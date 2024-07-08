package logic

import (
	"context"
	"fmt"
	"time"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/6tail/lunar-go/calendar"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChinaCalendarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChinaCalendarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChinaCalendarLogic {
	return &ChinaCalendarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChinaCalendarLogic) ChinaCalendar() (resp *types.ChinaCalendarReply, err error) {
	d := calendar.NewLunarFromDate(time.Now())

	var res types.ChinaCalendarReply
	// 宜
	s := d.GetDayYi()
	for i := s.Front(); i != nil; i = i.Next() {
		if len(res.Suitable) < 5 {
			res.Suitable = append(res.Suitable, i.Value.(string))
		}
	}
	// 忌(流派2)
	s = d.GetDayJiBySect(2)
	for i := s.Front(); i != nil; i = i.Next() {
		if len(res.Taboo) < 5 {
			res.Taboo = append(res.Taboo, i.Value.(string))
		}
	}
	solar := calendar.NewSolarFromDate(time.Now())
	res.Date = solar.String()
	res.ChinaDate = fmt.Sprintf("%s月%s", d.GetMonthInChinese(), d.GetDayInChinese())
	return &res, nil
}
