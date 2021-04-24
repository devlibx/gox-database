package gox_scylla

import (
	"context"
	"github.com/gocql/gocql"
	_ "github.com/gocql/gocql"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/metrics"
	goxdb "github.com/devlibx/gox-database"
)

type scyllaExecuteOp struct {
	*gocql.Session
	gox.CrossFunction
}

func (m *scyllaExecuteOp) ExecuteContext(ctx context.Context, metric metrics.LabeledMetric, query string, args ...interface{}) (int, error) {
	m.Counter("scylla_execute").Inc()
	var q *gocql.Query
	if args == nil {
		q = m.Session.Query(query)
	} else {
		q = m.Session.Query(query, args...)
	}
	if err := q.Exec(); err != nil {
		m.Counter("scylla_execute_error").Inc()
		return 0, &goxdb.DatabaseError{Op: goxdb.Execute, Query: query, Args: args, Err: err}
	}
	return 0, nil
}

func (m *scyllaExecuteOp) Execute(metric metrics.LabeledMetric, query string, args ...interface{}) (int, error) {
	return m.ExecuteContext(context.TODO(), metric, query, args...)
}
