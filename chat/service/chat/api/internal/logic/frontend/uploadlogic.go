package frontend

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/util/upload"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func CheckUploadImages(r *http.Request) bool {
	file, _, err := r.FormFile("file")
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	mimeType := http.DetectContentType(buffer)

	if mimeType != "image/jpeg" && mimeType != "image/png" {
		return false
	}
	return true
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadRequest, r *http.Request) (resp *types.UploadResponse, err error) {

	// 1. parse input , type multipart/form-data
	err = r.ParseMultipartForm(1000)
	if err != nil {
		return
	}
	// 2. retrieve file from posted form-data
	file, handler, err := r.FormFile("file")
	if err != nil {
		if err == http.ErrMissingFile {
			err = util.ReturnError(xerr.FileMiss)
			return
		}
		l.Logger.Errorf("Error retrieving file from form-data %v", err)
		return
	}
	defer file.Close()

	res, _, err := upload.NewOss(l.ctx, l.svcCtx, req.FileType).UploadFile(handler)
	if err != nil {
		l.Logger.Errorf("Error upload file from form-data %v", err)
		return
	}
	return &types.UploadResponse{
		File: res,
	}, nil
}
