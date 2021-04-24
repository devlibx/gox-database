module github.com/devlibx/gox-database/scylla

go 1.16

replace github.com/devlibx/gox-database => ../

require (
	github.com/devlibx/gox-base v0.0.36
	github.com/devlibx/gox-database v0.0.0-00010101000000-000000000000
	github.com/gocql/gocql v0.0.0-20210413161705-87a5d7a5ff74
)
