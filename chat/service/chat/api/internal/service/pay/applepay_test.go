package pay_test

import (
	"chat/common/redis"
	"chat/service/chat/api/internal/config"
	"chat/service/chat/api/internal/service/pay"
	"chat/service/chat/api/internal/svc"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
)

func TestApplepay_GetLastTransaction(t *testing.T) {
	ctx := context.Background()
	var c config.Config
	var configFile = flag.String("f", "../../../etc/chat-api.yaml", "the config file")
	conf.MustLoad(*configFile, &c)
	svc := svc.NewServiceContext(c)
	redis.Init(c.RedisCache[0].Host, c.RedisCache[0].Pass)

	appleyPayService := pay.NewApplePayService(ctx, svc)
	rsp, err := appleyPayService.GetLastTransactions("2000000535779432")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rsp)
}
