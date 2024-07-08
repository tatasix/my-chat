package repository

import (
	"chat/common/redis"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRoomRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRoomRepository(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomRepository {
	return &ChatRoomRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatRoomRepository) GetByChannelId(channelId int64) (ChatRoomPo *model.ChatRoom, err error) {
	return l.svcCtx.ChatRoomModel.FindOneByQuery(l.ctx,
		l.svcCtx.ChatRoomModel.RowBuilder().Where(squirrel.Eq{"channel_id": channelId}),
	)
}

func (l *ChatRoomRepository) Insert(ChatRoom *model.ChatRoom) (lastId int64, err error) {
	l.svcCtx.ChatRoomModel.BuildFiled(nil, ChatRoom)
	re, err := l.svcCtx.ChatRoomModel.Insert(l.ctx, ChatRoom)
	if err != nil {
		return
	}
	lastId, err = re.LastInsertId()
	return
}

func (l *ChatRoomRepository) Update(old, ChatRoom *model.ChatRoom) error {
	l.svcCtx.ChatRoomModel.BuildFiled(old, ChatRoom)
	return l.svcCtx.ChatRoomModel.Update(l.ctx, ChatRoom)
}

func (l *ChatRoomRepository) GetByChannelIdUseCache(channelId int64) (ChatRoomPo *model.ChatRoom, err error) {

	InfoCache, err1 := redis.Rdb.Get(l.ctx, fmt.Sprintf(redis.GetByChannelIdCacheKey, channelId)).Result()

	if err1 == nil {
		_ = json.Unmarshal([]byte(InfoCache), &ChatRoomPo)
		return
	} else {

		ChatRoomPo, err = l.GetByChannelId(channelId)
		if err == nil {
			redisData, err := json.Marshal(ChatRoomPo)
			if err == nil {
				redis.Rdb.Set(l.ctx, fmt.Sprintf(redis.GetByChannelIdCacheKey, channelId), string(redisData), redis.GetByChannelIdCacheExpireTime)
			}
		}
	}

	return
}

func (l *ChatRoomRepository) Delete(id int64) (err error) {
	return l.svcCtx.ChatRoomModel.SoftDelete(l.ctx, id)
}
