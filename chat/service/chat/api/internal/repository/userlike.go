package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLikeRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLikeRepository(ctx context.Context, svcCtx *svc.ServiceContext) *UserLikeRepository {
	return &UserLikeRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLikeRepository) GetByUserAndChatRecordId(user, chatRecordId string) (customerPo *model.UserLike, err error) {

	customerPo, err = l.svcCtx.UserLikeModel.FindOneByQuery(context.Background(),
		l.svcCtx.UserLikeModel.RowBuilder().Where(squirrel.Eq{"user": user}).Where(squirrel.Eq{"chat_record_id": chatRecordId}),
	)
	return
}

func (l *UserLikeRepository) GetByUser(user string) (customerPo []*model.UserLike, err error) {

	customerPo, err = l.svcCtx.UserLikeModel.FindAll(context.Background(),
		l.svcCtx.UserLikeModel.RowBuilder().Where(squirrel.Eq{"user": user}),
	)
	return
}

func (l *UserLikeRepository) GetById(id int64) (userLike *model.UserLike, err error) {
	return l.svcCtx.UserLikeModel.FindOne(context.Background(), id)
}

func (l *UserLikeRepository) Insert(userLike *model.UserLike) (sql.Result, error) {
	l.svcCtx.UserLikeModel.BuildFiled(nil, userLike)
	return l.svcCtx.UserLikeModel.Insert(l.ctx, userLike)
}

func (l *UserLikeRepository) DeleteByUserAndRecordId(user, chatRecordId string) (err error) {
	userLikeData, err := l.GetByUserAndChatRecordId(user, chatRecordId)
	if err != nil {
		return errors.New("数据不存在")
	}
	err = l.svcCtx.UserLikeModel.Delete(l.ctx, userLikeData.Id)
	return
}

func (l *UserLikeRepository) GetAll(chatRecordUser, startTime, endTime, order string, page, limit uint64) (chatRecordPos []*model.UserLike, count int64, err error) {
	l.Logger.Info(" GetUserLikeRecord UserLikeRepository start ")

	countBuilder := l.svcCtx.UserLikeModel.CountBuilder("id")
	rowBuilder := l.svcCtx.UserLikeModel.RowBuilder()
	if chatRecordUser != "" {
		countBuilder = countBuilder.Where(squirrel.Eq{"user": chatRecordUser})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"user": chatRecordUser})
	}

	if startTime != "" {
		countBuilder = countBuilder.Where("created_at >= ?", startTime)
		rowBuilder = rowBuilder.Where("created_at >= ?", startTime)
	}

	if endTime != "" {
		countBuilder = countBuilder.Where("created_at < ?", endTime)
		rowBuilder = rowBuilder.Where("created_at < ?", endTime)
	}

	count, err = l.svcCtx.UserLikeModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return
	}
	if count <= 0 {
		return nil, 0, nil
	}
	if order != "" {
		rowBuilder = rowBuilder.OrderBy(order)
	}
	if limit != 0 {
		offset := (page - 1) * limit
		rowBuilder = rowBuilder.Limit(limit).Offset(offset)
	}
	chatRecordPos, err = l.svcCtx.UserLikeModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}

	l.Logger.Info(" GetUserLikeRecord  end ")

	return
}
