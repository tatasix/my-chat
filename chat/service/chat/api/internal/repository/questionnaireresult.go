package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

const QuestionnaireResultStatusEnd = 100
const QuestionnaireResultSummary = 100
const QuestionnaireResultTypeGeneral = 1
const QuestionnaireResultTypeMBTI = 2

type QuestionnaireResultRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuestionnaireResultRepository(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionnaireResultRepository {
	return &QuestionnaireResultRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuestionnaireResultRepository) GetLastResult(kfId, user string) (customerPo *model.QuestionnaireResult, err error) {

	customerPo, err = l.svcCtx.QuestionnaireResultModel.FindOneByQuery(l.ctx,
		l.svcCtx.QuestionnaireResultModel.RowBuilder().Where(squirrel.Eq{"open_kf_id": kfId, "user": user}).OrderBy("id desc"),
	)
	return
}

func (l *QuestionnaireResultRepository) GetByRelationId(relationId int64) (customerPo []*model.QuestionnaireResult, err error) {

	customerPo, err = l.svcCtx.QuestionnaireResultModel.FindAll(l.ctx,
		l.svcCtx.QuestionnaireResultModel.RowBuilder().Where(squirrel.Eq{"relation_id": relationId}),
	)
	return
}

func (l *QuestionnaireResultRepository) Insert(QuestionnaireResult *model.QuestionnaireResult) (err error) {
	_, err = l.svcCtx.QuestionnaireResultModel.Insert(l.ctx, QuestionnaireResult)
	return
}

func (l *QuestionnaireResultRepository) Update(QuestionnaireResult *model.QuestionnaireResult) (err error) {
	return l.svcCtx.QuestionnaireResultModel.Update(l.ctx, QuestionnaireResult)
}

func (l *QuestionnaireResultRepository) GetByRelationIdAndType(relationId, scoreType int64) (customerPo *model.QuestionnaireResult, err error) {

	customerPo, err = l.svcCtx.QuestionnaireResultModel.FindOneByQuery(l.ctx,
		l.svcCtx.QuestionnaireResultModel.RowBuilder().Where(squirrel.Eq{"relation_id": relationId, "score_type": scoreType}),
	)
	return
}

func (l *QuestionnaireResultRepository) CountByMbti(mbti string) (mbtiNum, mbtiAll int64, err error) {
	rowBuilder := l.svcCtx.QuestionnaireResultModel.CountBuilder("mbti").Where(squirrel.Eq{"is_deleted": 0}).Where(squirrel.Eq{"mbti": mbti})
	countBuilder := l.svcCtx.QuestionnaireResultModel.CountBuilder("mbti").Where(squirrel.Eq{"is_deleted": 0}).Where(squirrel.NotEq{"mbti": ""})

	mbtiNum, err = l.svcCtx.QuestionnaireResultModel.FindCount(l.ctx, rowBuilder)
	mbtiAll, err = l.svcCtx.QuestionnaireResultModel.FindCount(l.ctx, countBuilder)

	return
}

func (l *QuestionnaireResultRepository) GetByRelationIdAndTypes(relationId int64, scoreType []int64) (customerPo *model.QuestionnaireResult, err error) {

	customerPo, err = l.svcCtx.QuestionnaireResultModel.FindOneByQuery(l.ctx,
		l.svcCtx.QuestionnaireResultModel.RowBuilder().Where(squirrel.Eq{"relation_id": relationId, "score_type": scoreType}),
	)
	return
}

func (l *QuestionnaireResultRepository) GetOne(relationId, scoreType int64, kfId, user string) (customerPo *model.QuestionnaireResult, err error) {

	rowBuilder := l.svcCtx.QuestionnaireResultModel.RowBuilder().OrderBy("id desc")
	if user != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"user": user})
	}
	if kfId != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"open_kf_id": kfId})
	}
	if relationId != 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"relation_id": relationId})
	}
	if scoreType != 0 {
		if scoreType == -1 {
			scoreType = 0
		}
		rowBuilder = rowBuilder.Where(squirrel.Eq{"score_type": scoreType})
	}
	customerPo, err = l.svcCtx.QuestionnaireResultModel.FindOneByQuery(l.ctx, rowBuilder)
	return
}
