package frontend

import (
	"chat/common/util"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/util/upload"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"os"
	"strconv"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadRequest, r *http.Request, w http.ResponseWriter) {

	path := upload.NewOss(l.ctx, l.svcCtx, req.FileType).GetFilePath()

	filePath := path + "/" + req.File
	l.Logger.Info("full file path: " + filePath)
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()
	//
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	// 设置响应头部
	util.SetHeader(w, req.File)
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// 将图片内容写入HTTP响应
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to write image content", http.StatusInternalServerError)
		return
	}

	return
}
