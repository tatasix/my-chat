package repository

import (
	"chat/common/redis"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

const ()

type PromptRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPromptRepository(ctx context.Context, svcCtx *svc.ServiceContext) *PromptRepository {
	return &PromptRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PromptRepository) GetByKfIdAndState(kfId string, state int32) (PromptPo *model.Prompt, err error) {
	return l.svcCtx.PromptModel.FindOneByQuery(l.ctx,
		l.svcCtx.PromptModel.RowBuilder().Where(squirrel.Eq{"kf_id": kfId}).Where(squirrel.Eq{"state_id": state}),
	)
}

func (l *PromptRepository) GetByKfId(kfId string) (PromptPo []*model.Prompt, err error) {
	return l.svcCtx.PromptModel.FindAll(l.ctx,
		l.svcCtx.PromptModel.RowBuilder().Where(squirrel.Eq{"kf_id": kfId}),
	)
}

func (l *PromptRepository) GetById(id int64) (Prompt *model.Prompt, err error) {
	return l.svcCtx.PromptModel.FindOne(context.Background(), id)
}

func (l *PromptRepository) GetByKfIdAndStateUseCache(kfId string, state int32) (promptPo *model.Prompt, err error) {

	promptPoCache, err1 := redis.Rdb.Get(l.ctx, fmt.Sprintf(redis.PromptGetByKfIdAndStateCacheKey, kfId, state)).Result()

	if err1 == nil {
		_ = json.Unmarshal([]byte(promptPoCache), &promptPo)
		return
	} else {

		promptPo, err = l.GetByKfIdAndState(kfId, state)
		if err == nil && promptPo != nil && promptPo.Id > 0 {
			redisData, _ := json.Marshal(promptPo)
			redis.Rdb.Set(l.ctx, fmt.Sprintf(redis.PromptGetByKfIdAndStateCacheKey, kfId, state), string(redisData), redis.PromptGetByKfIdAndStateCacheExpireTime)
		}
	}

	return
}

func (l *PromptRepository) Insert(prompt *model.Prompt) (sql.Result, error) {
	l.svcCtx.PromptModel.BuildFiled(nil, prompt)
	return l.svcCtx.PromptModel.Insert(l.ctx, prompt)
}

func (l *PromptRepository) Update(id int64, prompt *model.Prompt) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return errors.New("record not find")
	}
	l.svcCtx.PromptModel.BuildFiled(old, prompt)
	return l.svcCtx.PromptModel.Update(context.Background(), prompt)
}

func (l *PromptRepository) GetCountByKfId(kfId string) (count int64, err error) {
	countBuilder := l.svcCtx.PromptModel.CountBuilder("id").Where(squirrel.Eq{"kf_id": kfId})
	count, err = l.svcCtx.PromptModel.FindCount(l.ctx, countBuilder)
	return count, err
}
