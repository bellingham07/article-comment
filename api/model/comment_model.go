package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel

		FindOneById(ctx context.Context, id string) (*Comment, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

func (c *customCommentModel) FindOneById(ctx context.Context, id string) (*Comment, error) {
	//oid, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	fmt.Println("errr", err)
	//	return nil, ErrInvalidObjectId
	//}

	var data Comment

	key := prefixCommentCacheKey + id

	err := c.conn.FindOne(ctx, key, &data, id)
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewCommentModel returns a model for the mongo.
func NewCommentModel(url, db, collection string, c cache.CacheConf) CommentModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customCommentModel{
		defaultCommentModel: newDefaultCommentModel(conn),
	}
}
