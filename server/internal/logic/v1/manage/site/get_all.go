package site

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site"
)

type GetAll struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetAll(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetAll {
	return &GetAll{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetAll) GetAll(req *types.GetAllRequest) (resp []types.GetAllResponse, err error) {
	sites, err := l.svcCtx.Model.ManageSite.FindByCondition(l.ctx, nil)
	if err != nil {
		return nil, err
	}

	var list []types.GetAllResponse
	for _, s := range sites {
		if s.Status == "1" {
			list = append(list, types.GetAllResponse{
				Uuid:          s.Uuid,
				SiteName:      s.SiteName,
				PriceCurrency: s.PriceCurrency,
			})
		}
	}

	return list, nil
}
