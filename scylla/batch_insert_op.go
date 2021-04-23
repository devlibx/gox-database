package gox_scylla

import (
	"context"
	"errors"
	"github.com/gocql/gocql"
	"github.com/harishb2k/gox-base"
	errors2 "github.com/harishb2k/gox-base/errors"
	"github.com/harishb2k/gox-base/metrics"
	goxdb "github.com/harishb2k/gox-database"
)

type scyllaBatch struct {
	session *gocql.Session
	batch   *gocql.Batch
}

func (b *scyllaBatch) Add(query string, values ...interface{}) error {
	b.batch.Query(query, values...)
	return nil
}

// Batch insert operation implementation on top of scylla
type scyllaInsertBatchOp struct {
	*gocql.Session
	gox.CrossFunction
}

// Builds a new batch to run
func (s *scyllaInsertBatchOp) NewBatch() (goxdb.Batch, error) {
	return &scyllaBatch{
		session: s.Session,
		batch:   s.Session.NewBatch(gocql.LoggedBatch),
	}, nil
}

// Persist
func (s *scyllaInsertBatchOp) PersistBatchContext(ctx context.Context, metric metrics.LabeledMetric, batch goxdb.Batch) (interface{}, error) {
	s.Counter("scylla_persist_batch").Inc()
	if internalBatch, ok := batch.(*scyllaBatch); !ok {
		s.Counter("scylla_persist_batch_error").Inc()
		return nil, errors.New("failed to execute batch - expected batch to be of type scyllaBatch")
	} else {
		if err := s.Session.ExecuteBatch(internalBatch.batch); err != nil {
			s.Counter("scylla_persist_batch_error").Inc()
			return nil, errors2.Wrap(err, "failed to execute batch")
		}
		return nil, nil
	}
}

func (s *scyllaInsertBatchOp) PersistBatch(metric metrics.LabeledMetric, batch goxdb.Batch) (interface{}, error) {
	return s.PersistBatchContext(context.TODO(), metric, batch)
}
