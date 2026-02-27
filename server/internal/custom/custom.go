package custom

import (
	"context"
	"encoding/json"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/global"
	"github.com/jzero-io/jzero-admin/server/internal/logic/v1/manage/site"
	"github.com/jzero-io/jzero-admin/server/internal/model"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site"
	menutypes "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/menu"
	"github.com/jzero-io/jzero/core/stores/condition"
)

const siteScoreSyncInterval = 5 * time.Minute

type Custom struct {
	syncCancel context.CancelFunc
}

func New() *Custom {
	return &Custom{}
}

// Init Please add custom logic here.
func (c *Custom) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// auto gen casbin rules
	if err := InitCasbinRule(ctx, global.ServiceContext.Model, global.ServiceContext.CasbinEnforcer); err != nil {
		return err
	}

	return nil
}

// Start Please add custom logic here.
func (c *Custom) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	c.syncCancel = cancel
	go runSiteScoreSync(ctx)
}

// Stop Please add shut down logic here.
func (c *Custom) Stop() {
	if c.syncCancel != nil {
		c.syncCancel()
	}
}

func runSiteScoreSync(ctx context.Context) {
	ticker := time.NewTicker(siteScoreSyncInterval)
	defer ticker.Stop()
	doSync := func() {
		syncCtx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		sites, err := global.ServiceContext.Model.ManageSite.FindByCondition(syncCtx, nil, condition.Condition{
			Field:    manage_site.Status,
			Operator: condition.Equal,
			Value:    "1",
		})
		if err != nil {
			logx.Errorf("site score sync: list sites failed: %v", err)
			return
		}
		for _, s := range sites {
			if s.DbHost == "" || s.DbName == "" {
				continue
			}
			remoteScore, err := site.GetRemoteSiteRemainingScore(syncCtx, s)
			if err != nil {
				logx.Errorf("site score sync: site %s (%s): %v", s.Uuid, s.SiteName, err)
				continue
			}
			if err := global.ServiceContext.Model.ManageSite.UpdateFieldsByCondition(syncCtx, nil, map[string]any{
				string(manage_site.RemainingScore): remoteScore,
			}, condition.Condition{
				Field: manage_site.Uuid, Operator: condition.Equal, Value: s.Uuid,
			}); err != nil {
				logx.Errorf("site score sync: site %s (%s) update local: %v", s.Uuid, s.SiteName, err)
			}
		}
	}
	doSync()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			doSync()
		}
	}
}

func InitCasbinRule(ctx context.Context, model model.Model, enforcer *casbin.Enforcer) error {
	// get all role
	allRoles, err := model.ManageRole.FindByCondition(ctx, nil)
	if err != nil {
		return err
	}

	allMenus, err := model.ManageMenu.FindByCondition(ctx, nil)
	if err != nil {
		return err
	}

	// get role menu
	allRoleMenus, err := model.ManageRoleMenu.FindByCondition(ctx, nil)
	if err != nil {
		return err
	}

	var casbinRules [][]string
	for _, v := range allRoles {
		for _, arm := range allRoleMenus {
			if v.Uuid == arm.RoleUuid {
				for _, am := range allMenus {
					if arm.MenuUuid == am.Uuid {
						var permissions []menutypes.Permission
						err = json.Unmarshal([]byte(am.Permissions), &permissions)
						if err == nil {
							for _, perm := range permissions {
								if perm.Code != "" {
									if hasPolicy, _ := enforcer.HasPolicy(v.Uuid, perm.Code); !hasPolicy {
										casbinRules = append(casbinRules, []string{v.Uuid, perm.Code})
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if len(casbinRules) > 0 {
		_, err = enforcer.AddPolicies(casbinRules)
		if err != nil {
			return err
		}
	}
	return nil
}
