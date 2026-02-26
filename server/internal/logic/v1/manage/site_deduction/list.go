package site_deduction

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site_deduction"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site_deduction"
)

type List struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewList(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *List {
	return &List{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *List) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	records, total, err := l.svcCtx.Model.ManageSiteDeduction.PageByCondition(l.ctx, nil, condition.NewChain().
		Page(req.Current, req.Size).
		OrderByDesc(manage_site_deduction.CreateTime).
		Equal(manage_site_deduction.SiteUuid, req.SiteUuid, condition.WithSkip(req.SiteUuid == "")).
		Equal(manage_site_deduction.DeductionStatus, req.DeductionStatus, condition.WithSkip(req.DeductionStatus == "")).
		Like(manage_site_deduction.SiteName, "%"+req.SiteName+"%", condition.WithSkip(req.SiteName == "")).
		Build()...)
	if err != nil {
		return nil, err
	}

	var list []types.ManageSiteDeduction
	for _, r := range records {
		list = append(list, types.ManageSiteDeduction{
			Uuid:            r.Uuid,
			SiteUuid:        r.SiteUuid,
			Score:           fmt.Sprintf("%.2f", r.Score),
			DeductionMethod: r.DeductionMethod,
			DeductionStatus: r.DeductionStatus,
			Remark:          r.Remark,
			SiteName:        r.SiteName,
			PriceCurrency:   r.PriceCurrency,
			OperatorUuid:    r.OperatorUuid,
			CreateTime:      r.CreateTime.Format(time.DateTime),
			UpdateTime:      r.UpdateTime.Format(time.DateTime),
		})
	}

	resp = &types.ListResponse{
		Records: list,
		PageResponse: types.PageResponse{
			Current: req.Current,
			Size:    req.Size,
			Total:   total,
		},
	}
	return
}
