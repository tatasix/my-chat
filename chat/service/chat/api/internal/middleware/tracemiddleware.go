package middleware

import (
	"chat/common/util"
	"context"
	"net/http"
)

type TraceMiddleware struct {
}

func NewTraceMiddleware() *TraceMiddleware {
	return &TraceMiddleware{}
}

func (m *TraceMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get("trace-id")
		if val == "" {
			val = util.GenerateSnowflakeString()
		}
		ctx := context.WithValue(r.Context(), "trace-id", val)
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
