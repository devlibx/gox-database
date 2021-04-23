package gox_scylla

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/harishb2k/gox-base"
	"github.com/harishb2k/gox-base/metrics"
	goxdb "github.com/harishb2k/gox-database"
)

type scyllaSelectOp struct {
	*gocql.Session
	gox.CrossFunction
}

func (m *scyllaSelectOp) Find(metric metrics.LabeledMetric, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return m.FindContext(context.TODO(), metric, query, args...)
}

func (m *scyllaSelectOp) FindOne(metric metrics.LabeledMetric, query string, args ...interface{}) (map[string]interface{}, error) {
	return m.FindOneContext(context.TODO(), metric, query, args...)
}

func (m *scyllaSelectOp) FindContext(ctx context.Context, metric metrics.LabeledMetric, query string, args ...interface{}) ([]map[string]interface{}, error) {
	var iter *gocql.Iter
	if len(args) > 0 {
		iter = m.Session.Query(query, args...).Iter()
	} else {
		iter = m.Session.Query(query).Iter()
	}
	rowValues := make(map[string]interface{})
	var data []map[string]interface{}
	for iter.MapScan(rowValues) {
		data = append(data, rowValues)
		rowValues = make(map[string]interface{})
	}
	if len(data) > 0 {
		return data, nil
	} else {
		return nil, &goxdb.NoRecordFoundError{DatabaseError: goxdb.DatabaseError{Op: goxdb.FindAll, Query: query, Args: args, Err: nil}}
	}
}

func (m *scyllaSelectOp) FindOneContext(ctx context.Context, metric metrics.LabeledMetric, query string, args ...interface{}) (map[string]interface{}, error) {
	var iter *gocql.Iter
	if len(args) > 0 {
		iter = m.Session.Query(query, args...).Iter()
	} else {
		iter = m.Session.Query(query).Iter()
	}
	rowValues := make(map[string]interface{})
	var data map[string]interface{}
	for iter.MapScan(rowValues) {
		data = rowValues
		break
	}
	if data != nil {
		return data, nil
	} else {
		return nil, &goxdb.NoRecordFoundError{DatabaseError: goxdb.DatabaseError{Op: goxdb.Find, Query: query, Args: args, Err: nil}}
	}
}
