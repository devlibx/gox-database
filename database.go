package gox_database

import (
	"context"
	"errors"
	"fmt"
	"github.com/harishb2k/gox-base/metrics"
)

type Op string

// All operations supported in database
const (
	Open    Op = "open"
	Close   Op = "close"
	Insert  Op = "insert"
	Find    Op = "find"
	FindAll Op = "find_all"
	Delete  Op = "delete"
	Update  Op = "update"
	Execute Op = "execute"
)

// Error to represent a fail to insert
type DatabaseError struct {
	Op    Op
	Query string
	Args  []interface{}
	Err   error
	error
}

func (de *DatabaseError) Error() string {
	if de.Err != nil {
		return fmt.Sprintf("Op=%s, query=[%s] err=%v", de.Op, de.Query, de.Err)
	} else {
		return fmt.Sprintf("Op=%s, query=[%s])", de.Op, de.Query)
	}
}

func (de *DatabaseError) Unwrap() error {
	return de.Err
}

// Error when we do not find any record
type NoRecordFoundError struct {
	DatabaseError
}

func IsNoRecordFoundError(err error) (*NoRecordFoundError, bool) {
	var er *NoRecordFoundError
	if errors.As(err, &er) {
		return er, true
	} else {
		return nil, false
	}
}

// Insert operation to provide persist functionality
type InsertOp interface {
	// Persist to DB
	PersistContext(ctx context.Context, metric metrics.LabeledMetric, query string, args ...interface{}) (interface{}, error)
	Persist(metric metrics.LabeledMetric, query string, args ...interface{}) (interface{}, error)
}

// Select operation to fetch single or all records based on the query
type SelectOp interface {

	// Get all results based on the request query
	FindContext(ctx context.Context, metric metrics.LabeledMetric, query string, args ...interface{}) ([]map[string]interface{}, error)
	Find(metric metrics.LabeledMetric, query string, args ...interface{}) ([]map[string]interface{}, error)

	// Get one result base on the request query
	FindOneContext(ctx context.Context, metric metrics.LabeledMetric, query string, args ...interface{}) (map[string]interface{}, error)
	FindOne(metric metrics.LabeledMetric, query string, args ...interface{}) (map[string]interface{}, error)
}

// Execute a operation e.g. delete, update etc
type ExecuteOp interface {
	// Execute a query e.g. delete or update
	ExecuteContext(ctx context.Context, metric metrics.LabeledMetric, query string, args ...interface{}) (int, error)
	Execute(metric metrics.LabeledMetric, query string, args ...interface{}) (int, error)
}

// Execute one or many records in a single batch
type InsertBatchOp interface {
	NewBatch() (Batch, error)
	PersistBatchContext(ctx context.Context, metric metrics.LabeledMetric, batch Batch) (interface{}, error)
	PersistBatch(metric metrics.LabeledMetric, batch Batch) (interface{}, error)
}

// Provides all database operations
type Db interface {
	InsertOp
	SelectOp
	ExecuteOp
	InsertBatchOp
}
