package tool

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/util/upload"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"os"
)

type BatchUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchUploadLogic {
	return &BatchUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchUploadLogic) BatchUpload() (resp *types.Response, err error) {

	// 遍历指定文件夹中的所有文件并上传
	files, err := ioutil.ReadDir("./temp-files/scale/mbti/")
	if err != nil {
		fmt.Println("Error reading folder:", err)
		os.Exit(1)
	}
	var allName []string
	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			if fileName == ".DS_Store" {
				continue
			}

			filePath := "./temp-files/scale/mbti/" + file.Name()

			res, _, errUpload := upload.NewOss(l.ctx, l.svcCtx, 4).UploadFileByString(filePath)
			if errUpload != nil {
				l.Logger.Errorf("Error upload file from form-data %v", err)
				err = errUpload
				return
			}
			allName = append(allName, res)

			// 返回文件链接
			fmt.Println(allName)
			//return
		}
	}
	j, _ := json.Marshal(allName)
	resp = &types.Response{Message: string(j)}
	return
}
