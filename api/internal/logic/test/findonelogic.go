package test

import (
	"article-comment/api/internal/svc"
	"article-comment/api/internal/types"
	"article-comment/api/model"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type FindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneLogic {
	return &FindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOneLogic) FindOne(req *types.T) (resp *types.T, err error) {
	var date []model.Comment
	err = l.svcCtx.Comment.Find(l.ctx, &date, bson.M{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for i, comment := range date {
		fmt.Println(i, comment.Nickname)
	}
	fmt.Println("123")
	return
}
