package test

import (
	"context"

	"article-comment/api/internal/svc"
	"article-comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookLogic {
	return &GetBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBookLogic) GetBook(req *types.T) (resp *types.T, err error) {
	// todo: add your logic here and delete this line

	return
}
