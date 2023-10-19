package model

import "github.com/zeromicro/go-zero/core/stores/mon"

const EdgeinstanceCollectionName = "edgeinstance"

var _ EdgeinstanceModel = (*customEdgeinstanceModel)(nil)

type (
	// EdgeinstanceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEdgeinstanceModel.
	EdgeinstanceModel interface {
		edgeinstanceModel
	}

	customEdgeinstanceModel struct {
		*defaultEdgeinstanceModel
	}
)

// NewEdgeinstanceModel returns a model for the mongo.
func NewEdgeinstanceModel(url, db string) EdgeinstanceModel {
	conn := mon.MustNewModel(url, db, EdgeinstanceCollectionName)
	return &customEdgeinstanceModel{
		defaultEdgeinstanceModel: newDefaultEdgeinstanceModel(conn),
	}
}
