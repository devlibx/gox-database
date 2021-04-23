package gox_database

import (
	"context"
	"github.com/harishb2k/gox-base/metrics"
)

type Batch interface {
	Add(query string, values ...interface{}) error
}

type noOpBatch struct {
}

func (n *noOpBatch) Add(query string, values ...interface{}) error {
	return nil
}

type noOpBatchOp struct {
}

func (n noOpBatchOp) NewBatch() (Batch, error) {
	return &noOpBatch{}, nil
}

func (n noOpBatchOp) PersistBatchContext(ctx context.Context, metric metrics.LabeledMetric, batch Batch) (interface{}, error) {
	return nil, nil
}

func (n noOpBatchOp) PersistBatch(metric metrics.LabeledMetric, batch Batch) (interface{}, error) {
	return nil, nil
}

func NewNoOpBatchOp() InsertBatchOp {
	return &noOpBatchOp{}
}
