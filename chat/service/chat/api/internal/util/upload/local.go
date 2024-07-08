package upload

import (
	"chat/common/util"
	"chat/service/chat/api/internal/vars"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"chat/service/chat/api/internal/svc"
)

type Local struct {
	logger   logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	pathType uint32
}

// UploadFile
// @object: *Local
// @description: 上传文件
// @param: file *multipart.FileHeader
// @return: string, string, error
func (l *Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = util.MD5(name)
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	pathExt := GetFilePath(l.pathType)
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(l.svcCtx.Config.Local.StorePath+pathExt, os.ModePerm)
	if mkdirErr != nil {
		l.logger.Errorf(" function os.MkdirAll() Filed %+v", mkdirErr)
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := l.svcCtx.Config.Local.StorePath + pathExt + "/" + filename
	//filepath := l.svcCtx.Config.Local.Path + pathExt + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		l.logger.Errorf("function file.Open() Filed %+v", openError)
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		l.logger.Errorf("function os.Create() Filed %+v", openError)

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		l.logger.Errorf("function  io.Copy()  Filed %+v", openError)

		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}

	return l.svcCtx.Config.Domain + vars.DownloadApi + "file=" + filename + "&file_type=" + strconv.Itoa(int(l.pathType)), filename, nil
}

// DeleteFile
// @object: *Local
// @description: 删除文件
// @param: key string
// @return: error
func (l *Local) DeleteFile(key string) error {
	p := l.svcCtx.Config.Local.StorePath + GetFilePath(l.pathType) + "/" + key
	if strings.Contains(p, l.svcCtx.Config.Local.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
func (l *Local) GetFilePath() string {
	return l.svcCtx.Config.Local.StorePath + GetFilePath(l.pathType)
}

func (l *Local) UploadFileByString(localFile string) (string, string, error) {
	return "", "", nil
}
