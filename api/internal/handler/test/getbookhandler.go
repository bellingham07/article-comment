package test

import (
	"net/http"

	"article-comment/api/internal/logic/test"
	"article-comment/api/internal/svc"
	"article-comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.T
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := test.NewGetBookLogic(r.Context(), svcCtx)
		resp, err := l.GetBook(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}