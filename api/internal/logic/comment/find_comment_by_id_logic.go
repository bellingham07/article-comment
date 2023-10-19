package comment

import (
	"article-comment/api/internal/common/errorx"
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"article-comment/api/internal/svc"
	"article-comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindCommentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentByIdLogic {
	return &FindCommentByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentByIdLogic) FindCommentById(req *types.FindCommentByIdReq) (resp *types.FindCommentByIdResp, err error) {
	var (
		com []*types.Comment
	)
	err = l.svcCtx.Comment.Find(l.ctx, &com, bson.M{"_id": req.Id})
	if err != nil {
		return nil, errorx.Internal(err, "find by id fail").WithMetadata(errorx.Metadata{"req": req})
	}
	resp = new(types.FindCommentByIdResp)
	resp.List = com
	return
}
