package site

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site"
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
	if err = l.svcCtx.Model.ManageSite.InsertV2(l.ctx, nil, &manage_site.ManageSite{
		Uuid:          uuid.New().String(),
		SiteName:      req.SiteName,
		SiteLogo:      req.SiteLogo,
		PriceCurrency: req.PriceCurrency,
		ScoreStatType: req.ScoreStatType,
		Status:        req.Status,
		Remark:        req.Remark,
		DbHost:        req.DbHost,
		DbUsername:    req.DbUsername,
		DbPassword:    req.DbPassword,
		DbPort:        req.DbPort,
		DbName:        req.DbName,
		JwtSecret:     req.JwtSecret,
		AdminUrl:      req.AdminUrl,
		AdminUsername:  req.AdminUsername,
		ContactName:   req.ContactName,
		ContactPhone:  req.ContactPhone,
		ContactEmail:  req.ContactEmail,
	}); err != nil {
		return nil, err
	}

	return
}
