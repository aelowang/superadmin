package site

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
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site"
)

type Topup struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewTopup(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Topup {
	return &Topup{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *Topup) Topup(req *types.TopupRequest) (resp *types.TopupResponse, err error) {
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

	topupUuid := uuid.New().String()
	var siteData *manage_site.ManageSite

	err = l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		var innerErr error
		siteData, innerErr = l.svcCtx.Model.ManageSite.FindOneByUuid(ctx, session, req.Uuid)
		if innerErr != nil {
			return fmt.Errorf("site not found: %w", innerErr)
		}

		if innerErr = l.svcCtx.Model.ManageSiteTopup.InsertV2(ctx, session, &manage_site_topup.ManageSiteTopup{
			Uuid:          topupUuid,
			SiteUuid:      req.Uuid,
			Score:         score,
			TopupMethod:   "manual",
			TopupStatus:   "success",
			Remark:        req.Remark,
			SiteName:      siteData.SiteName,
			PriceCurrency: siteData.PriceCurrency,
			OperatorUuid:  info.Uuid,
		}); innerErr != nil {
			return innerErr
		}

		return l.svcCtx.Model.ManageSite.UpdateFieldsByCondition(ctx, session, map[string]any{
			string(manage_site.RemainingScore): siteData.RemainingScore + score,
			string(manage_site.TotalTopup):     siteData.TotalTopup + score,
		}, condition.Condition{
			Field:    manage_site.Uuid,
			Operator: condition.Equal,
			Value:    req.Uuid,
		})
	})
	if err != nil {
		return nil, err
	}

	if err = UpdateRemoteSiteRemainingScore(l.ctx, siteData, score); err != nil {
		_ = l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			_ = l.svcCtx.Model.ManageSiteTopup.DeleteByCondition(ctx, session, condition.Condition{
				Field: manage_site_topup.Uuid, Operator: condition.Equal, Value: topupUuid,
			})
			return l.svcCtx.Model.ManageSite.UpdateFieldsByCondition(ctx, session, map[string]any{
				string(manage_site.RemainingScore): siteData.RemainingScore,
				string(manage_site.TotalTopup):     siteData.TotalTopup,
			}, condition.Condition{
				Field: manage_site.Uuid, Operator: condition.Equal, Value: req.Uuid,
			})
		})
		return nil, fmt.Errorf("sync to site db failed: %w", err)
	}

	return
}
