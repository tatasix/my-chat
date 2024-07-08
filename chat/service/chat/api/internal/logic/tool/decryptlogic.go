package tool

import (
	"chat/common/util"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecryptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDecryptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecryptLogic {
	return &DecryptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DecryptLogic) Decrypt(req *types.EncryptMobileRequest) (resp *types.EncryptMobileResponse, err error) {
	resp = &types.EncryptMobileResponse{}
	var data []types.EncryptMobileDataResponse
	if len(req.Mobile) > 0 {
		for _, v := range req.Mobile {
			data = append(data, types.EncryptMobileDataResponse{
				Plaintext:  util.Decrypt(v),
				Ciphertext: v,
			})
		}
	}
	resp.Data = data
	return
}
