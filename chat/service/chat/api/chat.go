package main

import (
	"chat/common/redis"
	"chat/service/chat/api/internal/cron"
	"chat/service/chat/api/internal/startup"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"io"
	"log"
	"net/http"

	"chat/common/accesslog"
	"chat/common/response"
	"chat/common/wecom"
	"chat/common/xerr"
	"chat/service/chat/api/internal/config"
	"chat/service/chat/api/internal/handler"
	"chat/service/chat/api/internal/svc"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chat-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf,
		rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
			bodyByte, _ := io.ReadAll(r.Body)
			accesslog.ToLog(r, bodyByte, -1)
			response.Response(r, w, nil, errors.Wrapf(xerr.NewErrCode(xerr.UNAUTHORIZED), "鉴权失败 %v", err))
			return
		}),
		rest.WithNotFoundHandler(&NotFoundHandler{}),
		rest.WithNotAllowedHandler(&MethodNotMatchHandler{}),
	)
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/MP_verify_EzYNFyq0VmyM0r7Y.txt",
		Handler: filehandler("./MP_verify_EzYNFyq0VmyM0r7Y.txt"),
	})
	defer server.Stop()

	redis.Init(c.RedisCache[0].Host, c.RedisCache[0].Pass)
	defer redis.Close()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	cron.RegisterCrontab(ctx)
	wecom.WeCom.RestPort = c.RestConf.Port
	wecom.WeCom.Port = c.WeCom.Port
	wecom.WeCom.DefaultAgentSecret = c.WeCom.DefaultAgentSecret
	wecom.WeCom.CorpID = c.WeCom.CorpID
	wecom.WeCom.CustomerServiceSecret = c.WeCom.CustomerServiceSecret
	wecom.WeCom.Token = c.WeCom.Token
	wecom.WeCom.EncodingAESKey = c.WeCom.EncodingAESKey
	wecom.WeCom.Auth.AccessSecret = c.Auth.AccessSecret
	wecom.WeCom.Auth.AccessExpire = c.Auth.AccessExpire
	wecom.WeCom.MultipleApplication = startup.InitConfig(ctx)

	go wecom.XmlServe()
	go func() {
		log.Println(http.ListenAndServe(":7060", nil))
	}()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// disable stat
	logx.DisableStat()
	//group := service.NewServiceGroup()
	//defer group.Stop()
	//group.Add(server)
	//group.Add(im.Server{Ctx: ctx})
	//group.Start()
	server.Start()
}

type NotFoundHandler struct{}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bodyByte, _ := io.ReadAll(r.Body)
	accesslog.ToLog(r, bodyByte, -1)
	response.Response(r, w, nil, errors.Wrapf(xerr.NewErrCode(xerr.RouteNotFound), "接口不存在"))
	return
}

type MethodNotMatchHandler struct{}

func (h *MethodNotMatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bodyByte, _ := io.ReadAll(r.Body)
	accesslog.ToLog(r, bodyByte, -1)
	response.Response(r, w, nil, errors.Wrapf(xerr.NewErrCode(xerr.RouteNotMatch), "请求方式错误"))
	return
}

func addHeaders(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-CSRF-Token, Authorization, AccessToken, X-Requested-With, Content-Type, Accept,Token")
}

// 处理函数,传入文件地址
func filehandler(filepath string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, filepath)
	}
}
