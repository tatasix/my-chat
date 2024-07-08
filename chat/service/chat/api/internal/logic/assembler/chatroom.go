package assembler

import (
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"strconv"
)

func POTODTOGetChatRoomById(po *model.ChatRoom) (dto *types.ChatRoomAddResponse) {
	if po == nil {
		return
	}
	dto = &types.ChatRoomAddResponse{
		Id:        strconv.Itoa(int(po.Id)),
		ChannelId: strconv.Itoa(int(po.ChannelId)),
		User:      po.User,
		CreatedAt: po.CreatedAt.Format(vars.TimeFormat),
		UpdatedAt: po.UpdatedAt.Format(vars.TimeFormat),
	}

	return
}
