package rights

import (
	"chat/service/chat/api/internal/service/pay"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	// is the URL when testing your app in the sandbox and while your application is in review
	urlSandbox = "https://sandbox.itunes.apple.com/verifyReceipt"
	// is the URL when your app is live in the App Store
	urlProd = "https://buy.itunes.apple.com/verifyReceipt"
)

type ApplePayConfirmLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	applePayService *pay.ApplePayService
}

func NewApplePayConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplePayConfirmLogic {
	return &ApplePayConfirmLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		applePayService: pay.NewApplePayService(ctx, svcCtx),
	}
}

func (l *ApplePayConfirmLogic) ApplePayConfirm(req types.ApplePayConfirmReq) error {
	return l.applePayService.ReceiptVerification(req)
}
