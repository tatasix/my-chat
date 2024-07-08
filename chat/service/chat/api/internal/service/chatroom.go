package service

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRoomService struct {
	logx.Logger
	ctx                     context.Context
	svcCtx                  *svc.ServiceContext
	ChatRoomRepository      *repository.ChatRoomRepository
	ChatRoomUsersRepository *repository.ChatRoomUsersRepository
}

func NewChatRoomService(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomService {
	return &ChatRoomService{
		Logger:                  logx.WithContext(ctx),
		ctx:                     ctx,
		svcCtx:                  svcCtx,
		ChatRoomRepository:      repository.NewChatRoomRepository(ctx, svcCtx),
		ChatRoomUsersRepository: repository.NewChatRoomUsersRepository(ctx, svcCtx),
	}
}

func (l *ChatRoomService) Insert(ChatRoom *model.ChatRoom) (lastId int64, err error) {
	return l.ChatRoomRepository.Insert(ChatRoom)
}

func (l *ChatRoomService) Update(id int64, ChatRoom *model.ChatRoom) (err error) {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}
	l.svcCtx.ChatRoomModel.BuildFiled(old, ChatRoom)
	return l.svcCtx.ChatRoomModel.Update(context.Background(), ChatRoom)
}

func (l *ChatRoomService) Delete(channelId int64) (err error) {
	return l.ChatRoomRepository.Delete(channelId)
}

func (l *ChatRoomService) DeleteChatRoomUsersByChannelId(channelId int64) (err error) {
	return l.ChatRoomUsersRepository.Delete(channelId)

}

func (l *ChatRoomService) GetById(id int64) (chatRoom *model.ChatRoom, err error) {
	if chatRoom, err = l.svcCtx.ChatRoomModel.FindOne(context.Background(), id); nil != err && err != model.ErrNotFound {
		return
	}
	err = nil
	return
}
func (l *ChatRoomService) GetChatRoomByChannelId(channelId int64) (wechatUser *model.ChatRoom, err error) {
	return l.ChatRoomRepository.GetByChannelIdUseCache(channelId)
}

func (l *ChatRoomService) GetCountByChannel(channel int64) (int64, error) {
	return l.ChatRoomUsersRepository.GetCountByChannel(channel)
}

func (l *ChatRoomService) GetChatRoomUsers(channel int64) (wechatUser []*model.ChatRoomUsers, err error) {
	return l.ChatRoomUsersRepository.GetChatRoomUsers(channel)
}

func (l *ChatRoomService) AddChatRoomUsers(channel, isManager int64, user string) (err error) {
	return l.ChatRoomUsersRepository.Insert(&model.ChatRoomUsers{
		User:      user,
		ChannelId: channel,
		IsManager: isManager,
	})
}
