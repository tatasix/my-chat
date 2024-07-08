package middleware

import (
	"chat/common/jwt"
	"chat/common/response"
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/config"
	"net/http"
)

type LoginMiddleware struct {
	C config.Config
}

func NewLoginMiddleware(c config.Config) *LoginMiddleware {
	return &LoginMiddleware{
		C: c,
	}
}

func (s *LoginMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("authorization")
		if token == "" {
			response.Response(r, w, nil, util.ReturnError(xerr.InvalidToken))
			return
		}
		// 生成 token 并进行响应
		jwtObj := jwt.NewJwt(r.Context(), s.C.FrontendAuth.AccessSecret)

		// parseToken 解析token包含的信息
		user, err := jwtObj.ParseToken(token)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		query := r.URL.Query()

		// 设置参数值
		query.Set("user", user)

		// 更新URL中的参数
		r.URL.RawQuery = query.Encode()

		next(w, r)
	}
}
