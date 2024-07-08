package tool

import (
	"chat/common/util"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EncryptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEncryptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EncryptLogic {
	return &EncryptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EncryptLogic) Encrypt(req *types.EncryptMobileRequest) (resp *types.EncryptMobileResponse, err error) {
	resp = &types.EncryptMobileResponse{}
	var data []types.EncryptMobileDataResponse
	if len(req.Mobile) > 0 {
		for _, v := range req.Mobile {
			data = append(data, types.EncryptMobileDataResponse{
				Plaintext:  v,
				Ciphertext: util.Encrypt(v),
			})
		}
	}
	resp.Data = data
	return
}
