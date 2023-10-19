package comment

import (
	"article-comment/api/internal/common/errorx"
	"article-comment/api/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"

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
		coms []*model.Comment
	)
	err = l.svcCtx.Comment.Find(l.ctx, &coms, bson.M{})
	if err != nil {
		return nil, errorx.Internal(err, "查看全部评论失败").WithMetadata(errorx.Metadata{"req": req})
	}
	resp = new(types.FindAllCommentResp)
	resp.List = make([]*types.Comment, 0)
	for _, com := range coms {
		tmp := &types.Comment{
			ID:        com.ID,
			ArticleId: com.ArticleId,
			Content:   com.Content,
			UserId:    com.UserId,
			Nickname:  com.Nickname,
			LikeNum:   com.LikeNum,
			ReplyNum:  com.ReplyNum,
			State:     com.State,
			ParentId:  com.ParentId,
			UpdateAt:  strconv.FormatInt(com.UpdateAt.Unix(), 10),
			CreateAt:  strconv.FormatInt(com.CreateAt.Unix(), 10),
		}
		resp.List = append(resp.List, tmp)
	}
	return
}
