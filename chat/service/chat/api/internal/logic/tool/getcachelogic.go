package tool

import (
	"chat/common/redis"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCacheLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCacheLogic {
	return &GetCacheLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCacheLogic) GetCache(req *types.CacheRequest) (resp *types.CacheResponse, err error) {
	resp = &types.CacheResponse{
		Key: req.Key,
	}
	resp.Value, err = redis.Rdb.Get(l.ctx, req.Key).Result()

	return
}
