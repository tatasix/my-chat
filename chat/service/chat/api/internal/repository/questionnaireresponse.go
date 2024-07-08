package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type QuestionnaireResponseRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuestionnaireResponseRepository(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionnaireResponseRepository {
	return &QuestionnaireResponseRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuestionnaireResponseRepository) GetResponse(kfId, user string, relation int64) (QuestionnaireResponsePo []*model.QuestionnaireResponse, err error) {
	return l.svcCtx.QuestionnaireResponseModel.FindAll(l.ctx,
		l.svcCtx.QuestionnaireResponseModel.RowBuilder().Where(squirrel.Eq{"user": user, "open_kf_id": kfId, "relation_id": relation}),
	)
}

func (l *QuestionnaireResponseRepository) Insert(QuestionnaireResponse *model.QuestionnaireResponse) (lastId int64, err error) {
	res, err := l.svcCtx.QuestionnaireResponseModel.Insert(l.ctx, QuestionnaireResponse)
	if err != nil {
		return
	}
	lastId, err = res.LastInsertId()
	return
}

func (l *QuestionnaireResponseRepository) Update(QuestionnaireResponse *model.QuestionnaireResponse) error {
	return l.svcCtx.QuestionnaireResponseModel.Update(l.ctx, QuestionnaireResponse)
}

func (l *QuestionnaireResponseRepository) GetAll(kfId, user string, relation int64) (QuestionnaireResponsePo []*model.QuestionnaireResponse, err error) {
	rowBuilder := l.svcCtx.QuestionnaireResponseModel.RowBuilder()
	if kfId != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"open_kf_id": kfId})
	}
	if user != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"user": user})
	}
	if relation != 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"relation_id": relation})
	}
	return l.svcCtx.QuestionnaireResponseModel.FindAll(l.ctx, rowBuilder)
}

func (l *QuestionnaireResponseRepository) GetSum(relation, scoreType int64) (total int64, err error) {
	rowBuilder := l.svcCtx.QuestionnaireResponseModel.SumBuilder("score").Where(squirrel.Eq{"relation_id": relation})
	if scoreType != 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"score_type": scoreType})
	}
	return l.svcCtx.QuestionnaireResponseModel.FindSum(l.ctx, rowBuilder)
}

func (l *QuestionnaireResponseRepository) GetByRelationIdAndQuestionId(relationId, questionId int64) (QuestionnaireResponsePo *model.QuestionnaireResponse, err error) {
	return l.svcCtx.QuestionnaireResponseModel.FindOneByQuery(l.ctx,
		l.svcCtx.QuestionnaireResponseModel.RowBuilder().Where(squirrel.Eq{"relation_id": relationId, "question_id": questionId}),
	)
}

func (l *QuestionnaireResponseRepository) GetSumByTopics(kfId string, topics []int64, relationId int64) (total int64, err error) {
	rowBuilder := l.svcCtx.QuestionnaireResponseModel.SumBuilder("score").Where(squirrel.Eq{"question_id": topics, "open_kf_id": kfId, "relation_id": relationId})

	return l.svcCtx.QuestionnaireResponseModel.FindSum(l.ctx, rowBuilder)
}
