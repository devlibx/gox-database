package main

import (
	"fmt"
	"github.com/gocql/gocql"
	_ "github.com/gocql/gocql"
	"github.com/harishb2k/gox-base"
	"github.com/harishb2k/gox-base/metrics"
	goxdb "github.com/harishb2k/gox-database"
	goxscylla "github.com/harishb2k/gox-database/scylla"
	"time"
)

func main() {

	db, err := goxscylla.NewScyllaDb(&goxdb.Config{
		User:             "root",
		Password:         "root",
		Url:              []string{"localhost:9042"},
		Port:             9042,
		Db:               "example",
		CustomParameters: nil,
	}, gox.NewNoOpCrossFunction())
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	batch, _ := db.NewBatch()
	_ = batch.Add("insert into clients(client_id, tenant, component_name, client_internal_id, info, secrets) VALUES (?, ?, ?, ?, ?, ?);",
		gocql.UUIDFromTime(time.Now()),
		"payments",
		"workflow_1",
		gocql.UUIDFromTime(time.Now()),
		"",
		[]string{"password"},
	)
	_ = batch.Add("insert into clients(client_id, tenant, component_name, client_internal_id, info, secrets) VALUES (?, ?, ?, ?, ?, ?);",
		gocql.UUIDFromTime(time.Now()),
		"payments",
		"workflow_2",
		gocql.UUIDFromTime(time.Now()),
		"",
		[]string{"password"},
	)
	_ = batch.Add("insert into clients(client_id, tenant, component_name, client_internal_id, info, secrets) VALUES (?, ?, ?, ?, ?, ?);",
		gocql.UUIDFromTime(time.Now()),
		"payments",
		"workflow_3",
		gocql.UUIDFromTime(time.Now()),
		"",
		[]string{"password"},
	)
	_, err = db.PersistBatch(metrics.LabeledMetric{Name: "Batch"}, batch)
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Execute(metrics.LabeledMetric{Name: "Batch"}, "TRUNCATE TABLE job_by_job_id;")
	if err != nil {
		fmt.Println(err)
		panic("Failed to run execute")
	}
}
