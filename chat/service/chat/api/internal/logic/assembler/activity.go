package assembler

import (
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
)

func POTODTOGetStatistics(pos []*model.Statistics) (dto []types.Statistics) {
	if len(pos) <= 0 {
		return
	}

	for _, v := range pos {

		dto = append(dto, types.Statistics{
			Id:             v.Id,
			Date:           v.Date.Format(vars.TimeFormat2),
			DailyActive:    v.DailyActive,
			SevenActive:    v.SevenActive,
			FifteenActive:  v.FifteenActive,
			MonthlyActive:  v.MonthlyActive,
			TotalVisitor:   v.TotalVisitor,
			RegisteredUser: v.RegisteredUser,
			AddRegister:    v.AddRegister,
			AddVisitor:     v.AddVisitor,
			CreatedAt:      v.CreatedAt.Format(vars.TimeFormat),
			UpdatedAt:      v.UpdatedAt.Format(vars.TimeFormat),
		})
	}
	return
}
