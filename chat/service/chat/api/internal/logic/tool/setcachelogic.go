package tool

import (
	"chat/common/redis"
	"context"
	"time"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCacheLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCacheLogic {
	return &SetCacheLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetCacheLogic) SetCache(req *types.CacheRequest) (resp *types.CacheResponse, err error) {
	resp = &types.CacheResponse{
		Key:   req.Key,
		Value: req.Value,
	}
	_, err = redis.Rdb.Set(l.ctx, req.Key, req.Value, time.Duration(req.Expire)*time.Second).Result()

	return
}
