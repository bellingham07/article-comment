package comment

import (
	"article-comment/api/internal/common/errorx"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

	"article-comment/api/internal/svc"
	"article-comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.DeleteCommentResp, err error) {
	one, err := l.svcCtx.Comment.DeleteOne(l.ctx, bson.M{"_id": req.Id})
	fmt.Println(one)
	if err != nil {
		return nil, errorx.Internal(err, "delete fail").WithMetadata(errorx.Metadata{"req": req})
	}
	resp = new(types.DeleteCommentResp)
	resp.Message = "delete successful"
	return
}
