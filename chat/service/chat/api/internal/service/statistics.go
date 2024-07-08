package service

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type StatisticsService struct {
	logx.Logger
	ctx                   context.Context
	svcCtx                *svc.ServiceContext
	StatisticsRepository  *repository.StatisticsRepository
	ActivityLogRepository *repository.ActivityLogRepository
	wechatUserRepository  *repository.WechatUserRepository
}

func NewStatisticsService(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsService {
	return &StatisticsService{
		Logger:                logx.WithContext(ctx),
		ctx:                   ctx,
		svcCtx:                svcCtx,
		StatisticsRepository:  repository.NewStatisticsRepository(ctx, svcCtx),
		ActivityLogRepository: repository.NewActivityLogRepository(ctx, svcCtx),
		wechatUserRepository:  repository.NewWechatUserRepository(ctx, svcCtx),
	}
}

func (l *StatisticsService) Update(old, model *model.Statistics) (err error) {
	return l.StatisticsRepository.Update(old, model)
}

func (l *StatisticsService) GetAll(startCreatedAt, endCreatedAt string, page, pageSize uint64) (pos []*model.Statistics, count int64, err error) {
	return l.StatisticsRepository.GetAll(startCreatedAt, endCreatedAt, "id desc", page, pageSize)
}

func (l *StatisticsService) Statistics(ys string) (err error) {
	yesterday := time.Now().Add(-24 * time.Hour)

	if ys != "" {
		yesterday, err = time.Parse(vars.TimeFormat, ys)
		if err != nil {
			return
		}
	}

	yesterdayStart := yesterday.Format(vars.TimeFormat2) + " 00:00:00"
	yesterdayEnd := yesterday.Format(vars.TimeFormat2) + " 23:59:59"

	month := yesterday.AddDate(0, 0, -30)
	monthStart := month.Format(vars.TimeFormat2) + " 00:00:00"

	sevenDay := yesterday.AddDate(0, 0, -7)
	sevenStart := sevenDay.Format(vars.TimeFormat2) + " 00:00:00"

	fifteenDay := yesterday.AddDate(0, 0, -15)
	fifteenStart := fifteenDay.Format(vars.TimeFormat2) + " 00:00:00"

	//获取日活
	daily, err := l.ActivityLogRepository.GetDaily(yesterdayStart, yesterdayEnd)
	if err != nil {
		return
	}
	//获取月活
	monthly, err := l.ActivityLogRepository.GetDaily(monthStart, yesterdayEnd)
	if err != nil {
		return
	}
	//获取七天
	seven, err := l.ActivityLogRepository.GetDaily(sevenStart, yesterdayEnd)
	if err != nil {
		return
	}
	//获取十五天
	fifteen, err := l.ActivityLogRepository.GetDaily(fifteenStart, yesterdayEnd)
	if err != nil {
		return
	}
	//获取新增注册用户
	addRegister, err := l.wechatUserRepository.GetCountUser(yesterdayStart, yesterdayEnd, true)
	if err != nil {
		return
	}
	//获取新增游客数据
	addVisitor, err := l.wechatUserRepository.GetCountUser(yesterdayStart, yesterdayEnd, false)
	if err != nil {
		return
	}
	//获取总用户
	register, err := l.wechatUserRepository.GetCountUser("", yesterdayEnd, true)
	if err != nil {
		return
	}
	//获取总游客
	visitor, err := l.wechatUserRepository.GetCountUser("", yesterdayEnd, false)
	if err != nil {
		return
	}

	return l.StatisticsRepository.Insert(&model.Statistics{
		Date:           yesterday,
		DailyActive:    daily,
		SevenActive:    seven,
		FifteenActive:  fifteen,
		MonthlyActive:  monthly,
		TotalVisitor:   visitor,
		RegisteredUser: register,
		AddVisitor:     addVisitor,
		AddRegister:    addRegister,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
}
