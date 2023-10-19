package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ArticleId string             `bson:"article_id" json:"article_id"`
	Content   string             `bson:"content" json:"content"`
	UserId    string             `bson:"user_id" json:"user_id"`
	Nickname  string             `bson:"nickname" json:"nickname"`
	LikeNum   int64              `bson:"like_num" json:"like_num"`
	ReplyNum  int64              `bson:"reply_num" json:"reply_num"`
	State     string             `bson:"state" json:"state"`
	ParentId  string             `bson:"parent_id" json:"parent_id"`
	UpdateAt  time.Time          `bson:"update_at,omitempty" json:"update_at,omitempty"`
	CreateAt  time.Time          `bson:"create_at,omitempty" json:"create_at,omitempty"`
}
