package comment

import (
	"article-comment/api/internal/common/errorx"
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"article-comment/api/internal/svc"
	"article-comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindAllCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllCommentLogic {
	return &FindAllCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAllCommentLogic) FindAllComment(req *types.FindAllCommentReq) (resp *types.FindAllCommentResp, err error) {
	var (
		list []*types.Comment
	)
	err = l.svcCtx.Comment.Find(l.ctx, &list, bson.M{})
	if err != nil {
		return nil, errorx.Internal(err, "查看全部评论失败").WithMetadata(errorx.Metadata{"req": req})
	}
	resp.List = make([]*types.Comment, 0)
	return
}
