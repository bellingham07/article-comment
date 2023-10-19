package comment

import (
	"article-comment/api/internal/common/errorx"
	"article-comment/api/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"article-comment/api/internal/svc"
	"article-comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommentLogic) UpdateComment(req *types.UpdateCommentReq) (resp *types.UpdateCommentResp, err error) {
	update := bson.M{
		"$set": &model.Comment{
			Content:  "caojinbo is mvp",
			Nickname: "caoshuaishuai",
			UpdateAt: time.Now(),
		},
	}
	oid, err := primitive.ObjectIDFromHex(req.Id)
	id, err := l.svcCtx.Comment.UpdateByID(l.ctx, oid, update)
	fmt.Println(id)
	if err != nil {
		fmt.Println(err)
		return nil, errorx.Internal(err, "fail  to update").WithMetadata(errorx.Metadata{"req": req})
	}
	resp = new(types.UpdateCommentResp)
	resp.Message = "update successful"
	return
}
