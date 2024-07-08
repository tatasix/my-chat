package repository

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRoomUsersRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRoomUsersRepository(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomUsersRepository {
	return &ChatRoomUsersRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatRoomUsersRepository) GetByName(nickname string) (customerPo *model.ChatRoomUsers, err error) {

	customerPo, err = l.svcCtx.ChatRoomUsersModel.FindOneByQuery(l.ctx,
		l.svcCtx.ChatRoomUsersModel.RowBuilder().Where(squirrel.Eq{"nickname": nickname}),
	)
	return
}

func (l *ChatRoomUsersRepository) GetByUsers(users []string) (ChatRoomUsersPo []*model.ChatRoomUsers, err error) {
	return l.svcCtx.ChatRoomUsersModel.FindAll(l.ctx,
		l.svcCtx.ChatRoomUsersModel.RowBuilder().Where(squirrel.Eq{"user": users}),
	)
}

func (l *ChatRoomUsersRepository) GetByUser(user string) (ChatRoomUsersPo *model.ChatRoomUsers, err error) {
	return l.svcCtx.ChatRoomUsersModel.FindOneByQuery(l.ctx,
		l.svcCtx.ChatRoomUsersModel.RowBuilder().Where(squirrel.Eq{"user": user}),
	)
}

func (l *ChatRoomUsersRepository) UpdateById(id int64, ChatRoomUsers *model.ChatRoomUsers) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}
	l.svcCtx.ChatRoomUsersModel.BuildFiled(old, ChatRoomUsers)
	return l.svcCtx.ChatRoomUsersModel.Update(context.Background(), ChatRoomUsers)
}

func (l *ChatRoomUsersRepository) UpdateByUser(user string, ChatRoomUsers *model.ChatRoomUsers) error {
	old, err := l.GetByUser(user)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}
	l.svcCtx.ChatRoomUsersModel.BuildFiled(old, ChatRoomUsers)
	return l.svcCtx.ChatRoomUsersModel.Update(context.Background(), ChatRoomUsers)
}

func (l *ChatRoomUsersRepository) GetById(id int64) (ChatRoomUsers *model.ChatRoomUsers, err error) {
	return l.svcCtx.ChatRoomUsersModel.FindOne(context.Background(), id)
}

func (l *ChatRoomUsersRepository) GetByChannelId(channelId int64) (ChatRoomUsersPo *model.ChatRoomUsers, err error) {
	return l.svcCtx.ChatRoomUsersModel.FindOneByQuery(l.ctx,
		l.svcCtx.ChatRoomUsersModel.RowBuilder().Where(squirrel.Eq{"channel_id": channelId}),
	)
}

func (l *ChatRoomUsersRepository) Insert(ChatRoomUsers *model.ChatRoomUsers) (err error) {
	l.svcCtx.ChatRoomUsersModel.BuildFiled(nil, ChatRoomUsers)
	_, err = l.svcCtx.ChatRoomUsersModel.Insert(l.ctx, ChatRoomUsers)
	return
}

func (l *ChatRoomUsersRepository) Update(old, ChatRoomUsers *model.ChatRoomUsers) error {
	l.svcCtx.ChatRoomUsersModel.BuildFiled(old, ChatRoomUsers)
	return l.svcCtx.ChatRoomUsersModel.Update(l.ctx, ChatRoomUsers)
}

func (l *ChatRoomUsersRepository) GetCountByChannel(channelId int64) (int64, error) {

	countBuilder := l.svcCtx.ChatRoomUsersModel.CountBuilder("id").Where(squirrel.Eq{"channel_id": channelId})
	return l.svcCtx.ChatRoomUsersModel.FindCount(l.ctx, countBuilder)
}

func (l *ChatRoomUsersRepository) GetChatRoomUsers(channelId int64) (ChatRoomUsersPo []*model.ChatRoomUsers, err error) {
	return l.svcCtx.ChatRoomUsersModel.FindAll(l.ctx,
		l.svcCtx.ChatRoomUsersModel.RowBuilder().Where(squirrel.Eq{"channel_id": channelId}),
	)
}

func (l *ChatRoomUsersRepository) Delete(id int64) (err error) {
	return l.svcCtx.ChatRoomUsersModel.SoftDelete(l.ctx, id)
}
