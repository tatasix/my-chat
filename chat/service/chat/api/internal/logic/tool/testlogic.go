package tool

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/service/pay"
	"chat/service/chat/api/internal/service/questionnaire"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	wechatPayService *pay.WechatPayService
	alipayService    *pay.AlipayService
	RiskService      *service.RiskService
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		wechatPayService: pay.NewWechatPayService(ctx, svcCtx),
		alipayService:    pay.NewAlipayService(ctx, svcCtx),
		RiskService:      service.NewRiskService(ctx, svcCtx),
	}
}

func (l *TestLogic) Test() (resp *types.Response, err error) {
	res, err := repository.NewQuestionnaireQuestionRepository(l.ctx, l.svcCtx).GetByKfId("YTBNxzLNkCrpxZyCQkUmTCOa4Dk9QCXl")
	for _, v := range res {
		request := &model.QuestionnaireResponse{
			QuestionId: v.Id,
			RelationId: 6,
			MessageId:  "6",
			User:       "wmWpQ2GQAAW_bs_JzgqQjCsBt5Phjqjw",
			OpenKfId:   v.OpenKfId,
			Question:   v.Question,
			OptionId:   3,
			Answer:     "aaa",
		}
		questionnaire.NewQuestionnaireService(l.ctx, l.svcCtx).SaveQuestionnaire(request, v.Sort)
	}

	////charge := &pay.Charge{
	////	TradeNum:  util.GenerateSnowflakeString(),
	////	PayMethod: 1,
	////	MoneyFee:  1,
	////	Describe:  "test pay",
	////}
	////a, err := l.wechatPayService.GetQrCode(charge)
	////fmt.Printf("wechatPayService res:%s err:%+v", a, err)
	//
	////
	////b, err1 := l.alipayService.GetQrCode(charge)
	//l.RiskService.Reduce("wmWpQ2GQAAW_bs_JzgqQjCsBt5Phjqjw", "uqTIN13j6HKg2nYSyuTay6mHRQULNRSU", 1)
	return
}
