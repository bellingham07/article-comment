package httpc

import (
	"article-comment/api/internal/common/errorx"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go.uber.org/zap"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func RespSuccess(ctx context.Context, w http.ResponseWriter, resp interface{}) {
	var body Response
	body.Code = 200
	body.Msg = "success"
	body.Data = resp
	httpx.OkJsonCtx(ctx, w, body)
}
func RespError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	var (
		code     = http.StatusInternalServerError
		res      = Response{Code: code, Msg: "服务器开小差啦，稍后再来试一试"}
		metadata any
		appType  string
	)
	switch err.(type) {
	case *errorx.CodeError:
		customErr := errorx.From(err)
		res.Code = customErr.Code
		res.Msg = customErr.Msg // 	 展示给用户的友好错误信息(不暴露服务器内部信息)
		code = customErr.Code
		appType = customErr.Type
		metadata = customErr.Metadata
	}

	zap.L().Error(
		"", zap.Error(err), // 单独记录错误, 便于日志搜索; 不记录 Msg
		zap.Int("code", code),         // 记录状态码
		zap.String("type", appType),   // 记录业务type,便于定位错误
		zap.Any("metadata", metadata), // 记录元数据,入参
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
	)
	// 业务错误，http响应状态码为200，业务错误码在body中
	httpx.OkJsonCtx(ctx, w, res)
}
func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, &Response{401, "鉴权失败", nil})
}
