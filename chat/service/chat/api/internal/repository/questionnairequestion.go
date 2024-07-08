package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type QuestionnaireQuestionRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuestionnaireQuestionRepository(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionnaireQuestionRepository {
	return &QuestionnaireQuestionRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuestionnaireQuestionRepository) GetByKfId(kfId string) (QuestionnaireQuestionPo []*model.QuestionnaireQuestion, err error) {
	return l.svcCtx.QuestionnaireQuestionModel.FindAll(l.ctx,
		l.svcCtx.QuestionnaireQuestionModel.RowBuilder().Where(squirrel.Eq{"open_kf_id": kfId}),
	)
}

func (l *QuestionnaireQuestionRepository) GetById(id int64) (QuestionnaireQuestion *model.QuestionnaireQuestion, err error) {
	return l.svcCtx.QuestionnaireQuestionModel.FindOneByQuery(context.Background(), l.svcCtx.QuestionnaireQuestionModel.RowBuilder().Where(squirrel.Eq{"id": id}))
}

func (l *QuestionnaireQuestionRepository) Insert(QuestionnaireQuestion *model.QuestionnaireQuestion) (lastId int64, err error) {
	res, err := l.svcCtx.QuestionnaireQuestionModel.Insert(l.ctx, QuestionnaireQuestion)
	if err != nil {
		return
	}
	lastId, err = res.LastInsertId()
	return
}

func (l *QuestionnaireQuestionRepository) GetOption(options []model.QuestionnaireQuestionOption, optionId int64) (option model.QuestionnaireQuestionOption, err error) {
	if len(options) <= 0 {
		return
	}
	for _, v := range options {
		if v.Id == optionId {
			option = v
			return
		}
	}
	return
}
