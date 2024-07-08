package upload

import (
	"chat/service/chat/api/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
)

// OSS 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	GetFilePath() string
	DeleteFile(key string) error
	UploadFileByString(localFile string) (string, string, error)
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss(ctx context.Context, svcCtx *svc.ServiceContext, pathType uint32) OSS {
	switch svcCtx.Config.OssType {
	case "local":
		return &Local{logx.WithContext(ctx), ctx, svcCtx, pathType}
	case "qiniu":
		return &QiNiu{logx.WithContext(ctx), ctx, svcCtx, pathType}
	default:
		return &Local{logx.WithContext(ctx), ctx, svcCtx, pathType}
	}
}
