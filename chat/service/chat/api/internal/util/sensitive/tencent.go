package sensitive

import (
	"chat/service/chat/api/internal/svc"
	"context"
	"encoding/base64"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"net/url"
)

type Tencent struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (t *Tencent) Check(input string) (output bool) {
	bu, _ := url.Parse(t.svcCtx.Config.Sensitive.BucketURL) //"https://ding-1318601106.cos.ap-guangzhou.myqcloud.com"
	cu, _ := url.Parse(t.svcCtx.Config.Sensitive.CIURL)     //"https://ding-1318601106.ci.ap-guangzhou.myqcloud.com"
	b := &cos.BaseURL{BucketURL: bu, CIURL: cu}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.svcCtx.Config.Sensitive.SecretId,
			SecretKey: t.svcCtx.Config.Sensitive.SecretKey,
		},
	})
	opt := &cos.PutTextAuditingJobOptions{
		InputContent: base64.StdEncoding.EncodeToString([]byte(input)),
		Conf:         &cos.TextAuditingJobConf{},
	}
	res, _, err := c.CI.PutTextAuditingJob(context.Background(), opt)
	if err != nil {
		logx.WithContext(t.ctx).Errorf("TencentCheckError:%+v", err)
		return
	}
	if res == nil || res.JobsDetail == nil {
		return
	}
	return res.JobsDetail.Result == 0
}
