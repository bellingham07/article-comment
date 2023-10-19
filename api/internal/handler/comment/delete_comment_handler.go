package comment

import (
	"article-comment/api/internal/common/httpc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"article-comment/api/internal/logic/comment"
	"article-comment/api/internal/svc"
	"article-comment/api/internal/types"
)

func DeleteCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpc.RespError(r.Context(), w, r, err)
			return
		}
		l := comment.NewDeleteCommentLogic(r.Context(), svcCtx)
		resp, err := l.DeleteComment(&req)
		if err != nil {
			httpc.RespError(r.Context(), w, r, err)
		} else {
			httpc.RespSuccess(r.Context(), w, resp)
		}
	}
}
