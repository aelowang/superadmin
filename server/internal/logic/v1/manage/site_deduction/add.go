package site_deduction

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
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site_deduction"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site_deduction"
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

		if siteData.RemainingScore < score {
			return fmt.Errorf("insufficient remaining score: %.2f < %.2f", siteData.RemainingScore, score)
		}

		if err = l.svcCtx.Model.ManageSiteDeduction.InsertV2(ctx, session, &manage_site_deduction.ManageSiteDeduction{
			Uuid:            uuid.New().String(),
			SiteUuid:        req.SiteUuid,
			Score:           score,
			DeductionMethod: req.DeductionMethod,
			DeductionStatus: req.DeductionStatus,
			Remark:          req.Remark,
			SiteName:        siteData.SiteName,
			PriceCurrency:   siteData.PriceCurrency,
			OperatorUuid:    info.Uuid,
		}); err != nil {
			return err
		}

		return l.svcCtx.Model.ManageSite.UpdateFieldsByCondition(ctx, session, map[string]any{
			string(manage_site.RemainingScore): siteData.RemainingScore - score,
		}, condition.Condition{
			Field:    manage_site.Uuid,
			Operator: condition.Equal,
			Value:    req.SiteUuid,
		})
	})

	return
}
