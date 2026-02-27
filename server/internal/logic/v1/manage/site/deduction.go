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
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site_deduction"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site"
)

type Deduction struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewDeduction(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Deduction {
	return &Deduction{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *Deduction) Deduction(req *types.DeductionRequest) (resp *types.DeductionResponse, err error) {
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

	// 下分前先同步远程分数到本地
	siteData, err := l.svcCtx.Model.ManageSite.FindOneByUuid(l.ctx, nil, req.Uuid)
	if err != nil {
		return nil, fmt.Errorf("site not found: %w", err)
	}
	if siteData.DbHost != "" && siteData.DbName != "" {
		remoteScore, syncErr := GetRemoteSiteRemainingScore(l.ctx, siteData)
		if syncErr != nil {
			return nil, fmt.Errorf("sync remote score before deduction: %w", syncErr)
		}
		_ = l.svcCtx.Model.ManageSite.UpdateFieldsByCondition(l.ctx, nil, map[string]any{
			string(manage_site.RemainingScore): remoteScore,
		}, condition.Condition{
			Field: manage_site.Uuid, Operator: condition.Equal, Value: req.Uuid,
		})
		siteData.RemainingScore = remoteScore
	}
	if siteData.RemainingScore < score {
		return nil, fmt.Errorf("远程分数不够")
	}

	deductionUuid := uuid.New().String()

	err = l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		var innerErr error
		siteData, innerErr = l.svcCtx.Model.ManageSite.FindOneByUuid(ctx, session, req.Uuid)
		if innerErr != nil {
			return fmt.Errorf("site not found: %w", innerErr)
		}

		if innerErr = l.svcCtx.Model.ManageSiteDeduction.InsertV2(ctx, session, &manage_site_deduction.ManageSiteDeduction{
			Uuid:            deductionUuid,
			SiteUuid:        req.Uuid,
			Score:           score,
			DeductionMethod: "manual",
			DeductionStatus: "success",
			Remark:          req.Remark,
			SiteName:        siteData.SiteName,
			PriceCurrency:   siteData.PriceCurrency,
			OperatorUuid:    info.Uuid,
		}); innerErr != nil {
			return innerErr
		}

		return l.svcCtx.Model.ManageSite.UpdateFieldsByCondition(ctx, session, map[string]any{
			string(manage_site.RemainingScore): siteData.RemainingScore - score,
		}, condition.Condition{
			Field:    manage_site.Uuid,
			Operator: condition.Equal,
			Value:    req.Uuid,
		})
	})
	if err != nil {
		return nil, err
	}

	if err = UpdateRemoteSiteRemainingScore(l.ctx, siteData, -score); err != nil {
		_ = l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			_ = l.svcCtx.Model.ManageSiteDeduction.DeleteByCondition(ctx, session, condition.Condition{
				Field: manage_site_deduction.Uuid, Operator: condition.Equal, Value: deductionUuid,
			})
			return l.svcCtx.Model.ManageSite.UpdateFieldsByCondition(ctx, session, map[string]any{
				string(manage_site.RemainingScore): siteData.RemainingScore,
			}, condition.Condition{
				Field: manage_site.Uuid, Operator: condition.Equal, Value: req.Uuid,
			})
		})
		return nil, fmt.Errorf("sync to site db failed: %w", err)
	}

	return
}
