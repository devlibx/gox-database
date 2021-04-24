package gox_scylla

import (
	"github.com/gocql/gocql"
	_ "github.com/gocql/gocql"
	gox "github.com/devlibx/gox-base"
	goxdb "github.com/devlibx/gox-database"
)

type scyllaDb struct {
	scyllaInsertOp
	scyllaSelectOp
	scyllaExecuteOp
	scyllaInsertBatchOp
}

func NewScyllaDb(config *goxdb.Config, cf gox.CrossFunction) (goxdb.Db, error) {
	if config.Port == 0 {
		config.Port = 3306
	}
	if len(config.Url) == 0 {
		config.Url = []string{"127.0.0.1"}
	}
	cluster := gocql.NewCluster(config.Url...)
	cluster.Keyspace = config.Db

	_ = cf.RegisterCounter("scylla_persist", "scylla_persist", nil)
	_ = cf.RegisterCounter("scylla_persist_error", "scylla_persist_error", nil)
	_ = cf.RegisterCounter("scylla_persist_batch", "scylla_persist_batch", nil)
	_ = cf.RegisterCounter("scylla_persist_batch_error", "scylla_persist_batch_error", nil)
	_ = cf.RegisterCounter("scylla_find", "scylla_find", nil)
	_ = cf.RegisterCounter("scylla_find_error", "scylla_find_error", nil)
	_ = cf.RegisterCounter("scylla_find_one", "scylla_find_one", nil)
	_ = cf.RegisterCounter("scylla_find_one_error", "scylla_find_one_error", nil)
	_ = cf.RegisterCounter("scylla_execute", "scylla_execute", nil)
	_ = cf.RegisterCounter("scylla_execute_error", "scylla_execute_error", nil)

	if session, err := cluster.CreateSession(); err != nil {
		return nil, &goxdb.DatabaseError{Op: goxdb.Open, Query: "", Args: nil, Err: err}
	} else {
		return &scyllaDb{
			scyllaInsertOp:      scyllaInsertOp{Session: session, CrossFunction: cf},
			scyllaSelectOp:      scyllaSelectOp{Session: session, CrossFunction: cf},
			scyllaExecuteOp:     scyllaExecuteOp{Session: session, CrossFunction: cf},
			scyllaInsertBatchOp: scyllaInsertBatchOp{Session: session, CrossFunction: cf},
		}, nil
	}
}
