module main

go 1.16

require (
	github.com/ewkoll/aboutgo/schema v1.0.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
)

replace github.com/ewkoll/aboutgo/schema v1.0.0 => ./schema
