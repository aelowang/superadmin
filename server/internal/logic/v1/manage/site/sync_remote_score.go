package site

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_site"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/site"
)

type SyncRemoteScore struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewSyncRemoteScore(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *SyncRemoteScore {
	return &SyncRemoteScore{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *SyncRemoteScore) SyncRemoteScore(req *types.SyncRemoteScoreRequest) (err error) {
	siteData, err := l.svcCtx.Model.ManageSite.FindOneByUuid(l.ctx, nil, req.Uuid)
	if err != nil {
		return fmt.Errorf("site not found: %w", err)
	}
	remoteScore, err := GetRemoteSiteRemainingScore(l.ctx, siteData)
	if err != nil {
		return fmt.Errorf("get remote site score: %w", err)
	}
	return l.svcCtx.Model.ManageSite.UpdateFieldsByCondition(l.ctx, nil, map[string]any{
		string(manage_site.RemainingScore): remoteScore,
	}, condition.Condition{
		Field: manage_site.Uuid, Operator: condition.Equal, Value: req.Uuid,
	})
}
