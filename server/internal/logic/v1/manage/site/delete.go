package site

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site"
)

type Delete struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewDelete(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Delete {
	return &Delete{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *Delete) Delete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	if len(req.Uuids) == 0 {
		return nil, nil
	}

	err = l.svcCtx.Model.ManageSite.DeleteByCondition(l.ctx, nil, condition.Condition{
		Field:    manage_site.Uuid,
		Operator: condition.In,
		Value:    req.Uuids,
	})

	return
}
