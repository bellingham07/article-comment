package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ArticleId string             `bson:"article_id" json:"article_id"`
	Content   string             `bson:"content" json:"content"`
	Nickname  string             `bson:"nickname" json:"nickname"`
	LikeNum   int64              `bson:"like_num" json:"like_num"`
	ReplyNum  int64              `bson:"reply_num" json:"reply_num"`
	State     string             `bson:"state" json:"state"`
	ParentId  string             `bson:"parent_id" json:"parent_id"`
	UpdateAt  time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt  time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
