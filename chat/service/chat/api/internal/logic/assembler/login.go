package assembler

import (
	"chat/service/chat/api/internal/types"
)

func POTODTOGetShareConfig(appId, nonceStr, signature string, timestamp int64) (dto *types.GetShareConfigResponse) {
	dto = &types.GetShareConfigResponse{
		AppId:     appId,
		Timestamp: timestamp,
		NonceStr:  nonceStr,
		Signature: signature,
	}
	return
}
