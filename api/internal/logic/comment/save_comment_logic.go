package comment

import (
	"article-comment/api/internal/common/errorx"
	"article-comment/api/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"article-comment/api/internal/svc"
	"article-comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveCommentLogic {
	return &SaveCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveCommentLogic) SaveComment(req *types.SaveCommentReq) (resp *types.SaveCommentResp, err error) {
	com := &model.Comment{
		ArticleId: req.ArticleId,
		Content:   req.Content,
		UserId:    req.UserId,
		Nickname:  req.Nickname,
		LikeNum:   req.LikeNum,
		ReplyNum:  req.ReplyNum,
		State:     req.State,
		ParentId:  req.ParentId,
		CreateAt:  time.Now(),
	}
	one, err := l.svcCtx.Comment.InsertOne(l.ctx, &com)

	// 使用类型断言，调用Hex()函数，提取id关键字
	fmt.Println(one.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, errorx.Internal(err, "fail to insert").WithMetadata(errorx.Metadata{"req": req})
	}
	resp = new(types.SaveCommentResp)
	resp.Message = "insert successful"
	return
}
