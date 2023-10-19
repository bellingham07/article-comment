package comment

import (
	"article-comment/api/internal/common/errorx"
	"article-comment/api/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"

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
		com []*model.Comment
	)
	oid, err := primitive.ObjectIDFromHex(req.Id)
	err = l.svcCtx.Comment.Find(l.ctx, &com, bson.M{"_id": oid})
	if err != nil {
		return nil, errorx.Internal(err, "find by id fail").WithMetadata(errorx.Metadata{"req": req})
	}

	resp = new(types.FindCommentByIdResp)
	resp.List = make([]*types.Comment, 0)
	for _, comment := range com {
		tmp := &types.Comment{
			ID:        comment.ID,
			ArticleId: comment.ArticleId,
			Content:   comment.Content,
			UserId:    comment.UserId,
			Nickname:  comment.Nickname,
			LikeNum:   comment.LikeNum,
			ReplyNum:  comment.ReplyNum,
			State:     comment.State,
			ParentId:  comment.ParentId,
			UpdateAt:  strconv.FormatInt(comment.UpdateAt.Unix(), 10),
			CreateAt:  strconv.FormatInt(comment.CreateAt.Unix(), 10),
		}
		resp.List = append(resp.List, tmp)
	}
	return
}
