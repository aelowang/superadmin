package site

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site"
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
	sites, total, err := l.svcCtx.Model.ManageSite.PageByCondition(l.ctx, nil, condition.NewChain().
		Page(req.Current, req.Size).
		OrderByDesc(manage_site.CreateTime).
		Like(manage_site.SiteName, "%"+req.SiteName+"%", condition.WithSkip(req.SiteName == "")).
		Equal(manage_site.PriceCurrency, req.PriceCurrency, condition.WithSkip(req.PriceCurrency == "")).
		Equal(manage_site.Status, req.Status, condition.WithSkip(req.Status == "")).
		Build()...)
	if err != nil {
		return nil, err
	}

	var records []types.ManageSite
	for _, s := range sites {
		records = append(records, types.ManageSite{
			Uuid:           s.Uuid,
			SiteName:       s.SiteName,
			SiteLogo:       s.SiteLogo,
			PriceCurrency:  s.PriceCurrency,
			ScoreStatType:  s.ScoreStatType,
			Status:         s.Status,
			Remark:         s.Remark,
			DbHost:         s.DbHost,
			DbUsername:     s.DbUsername,
			DbPassword:     "******",
			DbPort:         s.DbPort,
			DbName:         s.DbName,
			JwtSecret:      "******",
			AdminUrl:       s.AdminUrl,
			AdminUsername:   s.AdminUsername,
			ContactName:    s.ContactName,
			ContactPhone:   s.ContactPhone,
			ContactEmail:   s.ContactEmail,
			RemainingScore: fmt.Sprintf("%.2f", s.RemainingScore),
			TotalTopup:     fmt.Sprintf("%.2f", s.TotalTopup),
			CreateTime:     s.CreateTime.Format(time.DateTime),
			UpdateTime:     s.UpdateTime.Format(time.DateTime),
		})
	}

	resp = &types.ListResponse{
		Records: records,
		PageResponse: types.PageResponse{
			Current: req.Current,
			Size:    req.Size,
			Total:   total,
		},
	}
	return
}
