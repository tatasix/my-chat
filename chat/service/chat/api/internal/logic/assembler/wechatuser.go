package assembler

import (
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
)

func POTODTOGetUser(wechatUserPo *model.WechatUser) (dto *types.GetUserResponse) {
	if wechatUserPo == nil {
		return
	}
	dto = &types.GetUserResponse{}
	dto.Id = wechatUserPo.Id
	dto.User = wechatUserPo.User
	dto.UserType = wechatUserPo.UserType
	dto.Nickname = wechatUserPo.Nickname
	dto.Mobile = wechatUserPo.Mobile
	dto.Level = wechatUserPo.Level
	dto.LevelExpire = wechatUserPo.LevelExpire
	dto.Name = wechatUserPo.Name
	dto.Status = wechatUserPo.Status
	dto.Avatar = wechatUserPo.Avatar
	dto.Gender = wechatUserPo.Gender
	dto.Unionid = wechatUserPo.Unionid
	dto.Openid = wechatUserPo.Openid
	dto.Province = wechatUserPo.Province
	dto.City = wechatUserPo.City
	dto.Country = wechatUserPo.Country
	dto.Mbti = wechatUserPo.Mbti
	dto.Constellation = wechatUserPo.Constellation
	dto.Birthday = wechatUserPo.Birthday
	dto.Constellation = wechatUserPo.Constellation
	dto.CreatedAt = wechatUserPo.CreatedAt.Format(vars.TimeFormat)
	dto.UpdatedAt = wechatUserPo.UpdatedAt.Format(vars.TimeFormat)
	return
}
