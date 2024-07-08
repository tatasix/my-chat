package tool

import (
	"chat/common/redis"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCacheLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCacheLogic {
	return &DeleteCacheLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCacheLogic) DeleteCache(req *types.CacheRequest) (resp *types.CacheResponse, err error) {
	resp = &types.CacheResponse{
		Key: req.Key,
	}
	err = redis.Rdb.Del(l.ctx, req.Key).Err()

	return
}
