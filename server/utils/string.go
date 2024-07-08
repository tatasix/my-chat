package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/vars"
	"time"
)

func TimeToString(info *time.Time) (res string) {
	if info != nil {
		res = info.String()
	}
	return
}

func TimeFormat(origin string) (res time.Time) {
	res, _ = time.Parse(vars.TimeFormat, origin)
	return
}
