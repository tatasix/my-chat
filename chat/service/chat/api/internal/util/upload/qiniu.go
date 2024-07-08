package upload

import (
	"chat/common/util"
	"chat/service/chat/api/internal/svc"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
	"os"
	"path"
	"time"
)

type QiNiu struct {
	logger   logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	pathType uint32
}

// UploadFile
// @description: 上传文件
// @param: file *multipart.FileHeader
// @return: string, string, error
func (q *QiNiu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: q.svcCtx.Config.QiNiu.Bucket}
	mac := qbox.NewMac(q.svcCtx.Config.QiNiu.AccessKey, q.svcCtx.Config.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiNiuConfig(q.svcCtx.Config.QiNiu.UseHTTPS, q.svcCtx.Config.QiNiu.UseCdnDomains, q.svcCtx.Config.QiNiu.Zone)
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		q.logger.Errorf(" function file.Open() Filed %+v", openError)
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close()
	// 创建文件 defer 关闭
	pathExt := GetFilePath(q.pathType)
	fileKey := fmt.Sprintf("%s/%s/%s%s", pathExt, time.Now().Format("20060102"), util.GenerateSnowflakeString(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		q.logger.Errorf(" function formUploader.Put() Filed %+v", putErr)

		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	return q.svcCtx.Config.QiNiu.ImgPath + "/" + ret.Key, ret.Key, nil
}

// DeleteFile
// @description: 删除文件
// @param: key string
// @return: error
func (q *QiNiu) DeleteFile(key string) error {
	mac := qbox.NewMac(q.svcCtx.Config.QiNiu.AccessKey, q.svcCtx.Config.QiNiu.SecretKey)
	cfg := qiNiuConfig(q.svcCtx.Config.QiNiu.UseHTTPS, q.svcCtx.Config.QiNiu.UseCdnDomains, q.svcCtx.Config.QiNiu.Zone)
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(q.svcCtx.Config.QiNiu.Bucket, key); err != nil {
		q.logger.Errorf(" function bucketManager.Delete() Filed %+v", err)

		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// qiNiuConfig
// @description: 根据配置文件进行返回七牛云的配置
// @return: *storage.Config
func qiNiuConfig(useHTTPS, useCdnDomains bool, zone string) *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      useHTTPS,
		UseCdnDomains: useCdnDomains,
	}
	switch zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}

func (q *QiNiu) GetFilePath() string {

	return ""
}

// UploadFileByString
// @description: 上传文件
// @param: localFile string
// @return: string, string, error
func (q *QiNiu) UploadFileByString(localFile string) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: q.svcCtx.Config.QiNiu.Bucket}
	mac := qbox.NewMac(q.svcCtx.Config.QiNiu.AccessKey, q.svcCtx.Config.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiNiuConfig(q.svcCtx.Config.QiNiu.UseHTTPS, q.svcCtx.Config.QiNiu.UseCdnDomains, q.svcCtx.Config.QiNiu.Zone)
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, err := os.Open(localFile)
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	// 创建文件 defer 关闭
	pathExt := GetFilePath(q.pathType)

	fileKey := fmt.Sprintf("%s/%s/%s%s", pathExt, time.Now().Format("20060102"), util.GenerateSnowflakeString(), path.Base(localFile)) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.PutFile(context.Background(), &ret, upToken, fileKey, localFile, &putExtra)
	if putErr != nil {
		q.logger.Errorf(" function formUploader.Put() Filed %+v", putErr)

		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	return q.svcCtx.Config.QiNiu.ImgPath + "/" + ret.Key, ret.Key, nil
}
