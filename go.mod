module main

go 1.16

require (
	github.com/ewkoll/aboutgo/schema v1.0.0
	github.com/ewkoll/aboutgo/server v1.0.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.2.0
	github.com/pkg/errors v0.9.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

replace github.com/ewkoll/aboutgo/schema v1.0.0 => ./schema

replace github.com/ewkoll/aboutgo/server v1.0.0 => ./server
