package site_topup

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/jzero-io/jzero-admin/core-engine/helper/auth"
	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site_topup"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site_topup"
)

type Add struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAdd(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Add {
	return &Add{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *Add) Add(req *types.AddRequest) (resp *types.AddResponse, err error) {
	score, err := strconv.ParseFloat(req.Score, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid score: %s", req.Score)
	}
	if score <= 0 {
		return nil, fmt.Errorf("score must be greater than 0")
	}

	info, err := auth.Info(l.ctx)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		siteData, err := l.svcCtx.Model.ManageSite.FindOneByUuid(ctx, session, req.SiteUuid)
		if err != nil {
			return fmt.Errorf("site not found: %w", err)
		}

		if err = l.svcCtx.Model.ManageSiteTopup.InsertV2(ctx, session, &manage_site_topup.ManageSiteTopup{
			Uuid:          uuid.New().String(),
			SiteUuid:      req.SiteUuid,
			Score:         score,
			TopupMethod:   req.TopupMethod,
			TopupStatus:   req.TopupStatus,
			Remark:        req.Remark,
			SiteName:      siteData.SiteName,
			PriceCurrency: siteData.PriceCurrency,
			OperatorUuid:  info.Uuid,
		}); err != nil {
			return err
		}

		return l.svcCtx.Model.ManageSite.UpdateFieldsByCondition(ctx, session, map[string]any{
			string(manage_site.RemainingScore): siteData.RemainingScore + score,
			string(manage_site.TotalTopup):     siteData.TotalTopup + score,
		}, condition.Condition{
			Field:    manage_site.Uuid,
			Operator: condition.Equal,
			Value:    req.SiteUuid,
		})
	})

	return
}
