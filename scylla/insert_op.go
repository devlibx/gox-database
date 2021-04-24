package gox_scylla

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/metrics"
)

type scyllaInsertOp struct {
	*gocql.Session
	gox.CrossFunction
}

func (s scyllaInsertOp) PersistContext(ctx context.Context, metric metrics.LabeledMetric, query string, args ...interface{}) (interface{}, error) {
	statement := s.Session.Query(query, args...)
	if err := statement.Exec(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s scyllaInsertOp) Persist(metric metrics.LabeledMetric, query string, args ...interface{}) (interface{}, error) {
	statement := s.Session.Query(query, args...)
	if err := statement.Exec(); err != nil {
		return nil, err
	}
	return nil, nil
}
