package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type QuestionnaireResultMbtiRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuestionnaireResultMbtiRepository(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionnaireResultMbtiRepository {
	return &QuestionnaireResultMbtiRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuestionnaireResultMbtiRepository) GetLastResult(kfId, user string) (po *model.QuestionnaireResultMbti, err error) {

	po, err = l.svcCtx.QuestionnaireResultMbtiModel.FindOneByQuery(l.ctx,
		l.svcCtx.QuestionnaireResultMbtiModel.RowBuilder().Where(squirrel.Eq{"open_kf_id": kfId, "user": user}).OrderBy("id desc"),
	)
	return
}

func (l *QuestionnaireResultMbtiRepository) GetByRelationId(relationId int64) (po *model.QuestionnaireResultMbti, err error) {

	po, err = l.svcCtx.QuestionnaireResultMbtiModel.FindOneByQuery(l.ctx,
		l.svcCtx.QuestionnaireResultMbtiModel.RowBuilder().Where(squirrel.Eq{"relation_id": relationId}),
	)
	return
}

func (l *QuestionnaireResultMbtiRepository) Insert(QuestionnaireResultMbti *model.QuestionnaireResultMbti) (id int64, err error) {
	res, err := l.svcCtx.QuestionnaireResultMbtiModel.Insert(l.ctx, QuestionnaireResultMbti)
	if err != nil {
		return
	}
	id, _ = res.LastInsertId()
	return
}

func (l *QuestionnaireResultMbtiRepository) Update(QuestionnaireResultMbti *model.QuestionnaireResultMbti) (err error) {
	return l.svcCtx.QuestionnaireResultMbtiModel.Update(l.ctx, QuestionnaireResultMbti)
}
