package site

import (
	"context"
	"net/http"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site"
)

type Edit struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewEdit(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Edit {
	return &Edit{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *Edit) Edit(req *types.EditRequest) (resp *types.EditResponse, err error) {
	data, err := l.svcCtx.Model.ManageSite.FindOneByUuid(l.ctx, nil, req.Uuid)
	if err != nil {
		return nil, err
	}

	newData := lo.FromPtr(data)
	newData.SiteName = req.SiteName
	newData.SiteLogo = req.SiteLogo
	newData.PriceCurrency = req.PriceCurrency
	newData.ScoreStatType = req.ScoreStatType
	newData.Status = req.Status
	newData.Remark = req.Remark
	newData.DbHost = req.DbHost
	newData.DbUsername = req.DbUsername
	newData.DbPassword = req.DbPassword
	newData.DbPort = req.DbPort
	newData.DbName = req.DbName
	newData.JwtSecret = req.JwtSecret
	newData.AdminUrl = req.AdminUrl
	newData.AdminUsername = req.AdminUsername
	newData.ContactName = req.ContactName
	newData.ContactPhone = req.ContactPhone
	newData.ContactEmail = req.ContactEmail

	if err = l.svcCtx.Model.ManageSite.Update(l.ctx, nil, lo.ToPtr(newData)); err != nil {
		return nil, err
	}

	return
}
