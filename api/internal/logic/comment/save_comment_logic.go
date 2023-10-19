package comment

import (
	"article-comment/api/internal/common/errorx"
	"article-comment/api/model"
	"context"
	"fmt"
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
		ArticleId: req.Com.ArticleId,
		Content:   req.Com.Content,
		UserId:    req.Com.UserId,
		Nickname:  req.Com.Nickname,
		LikeNum:   req.Com.LikeNum,
		ReplyNum:  req.Com.ReplyNum,
		State:     req.Com.State,
		ParentId:  req.Com.ParentId,
		UpdateAt:  time.Now(),
		CreateAt:  time.Now(),
	}
	one, err := l.svcCtx.Comment.InsertOne(l.ctx, &com)
	fmt.Println(one)
	if err != nil {
		return nil, errorx.Internal(err, "fail to insert").WithMetadata(errorx.Metadata{"req": req})
	}
	resp = new(types.SaveCommentResp)
	resp.Message = "insert successful"
	return
}
