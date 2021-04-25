package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func queryNoRow(db *sql.DB, ctx context.Context) error {

	ret, err := db.ExecContext(ctx, "update new_table set id = 100 where id = 1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)

	return nil
}

func main() {
	db := DB{driverName: "mysql", dsn: "demo:demo@tcp(127.0.0.1:3306)/oms"}
	err := db.Load()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = db.Exec(queryNoRow)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("Exec failed with no rows and %v", err)
	}
}
