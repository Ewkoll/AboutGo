package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/ewkoll/aboutgo/schema"
	"github.com/google/uuid"
	xerrors "github.com/pkg/errors"
)

func queryNoRow(db *sql.DB, ctx context.Context) error {
	var user schema.User
	err := db.QueryRowContext(ctx, "select * from db_users where id = '33'").Scan(&user)
	switch {
	case err == sql.ErrNoRows:
		return xerrors.Wrap(err, "main: query failed with no data")
	case err != nil:
		return xerrors.Wrap(err, "main: execute failed")
	}
	return nil
}

func doInsert(db *sql.DB, ctx context.Context) error {
	stmt, err := db.Prepare("insert into db_users(id, role_id, user_name) values(?, ?, ?)")
	if err != nil {
		return xerrors.Wrap(err, "main: prepare sql failed")
	}
	defer stmt.Close()

	user := schema.User{ID: uuid.New().String(), RoleID: "1", UserName: "Ewkoll"}
	ret, err := stmt.Exec(user.ID, user.RoleID, user.UserName)
	if err != nil {
		return xerrors.Wrap(err, "main: stmt execute sql failed")
	}

	rowCount, _ := ret.RowsAffected()
	fmt.Printf("affected row count = %d\n", rowCount)
	return nil
}

func doUpdate(db *sql.DB, ctx context.Context) error {
	var testSql = `select id from db_users limit ?;`
	rowCount := 1
	rows, err := db.Query(testSql, rowCount)
	if err != nil {
		return xerrors.Wrap(err, "main: sql query failed")
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return xerrors.Wrap(err, "main: sql query failed")
		}

		rowsUpdate, err := db.ExecContext(ctx, "update db_users set user_name = 'new_name' where id = ?;", id)
		if err != nil {
			return xerrors.Wrap(err, "main: sql update failed")
		}
		rowCount, _ := rowsUpdate.RowsAffected()
		fmt.Printf("affected row count = %d\n", rowCount)
	}

	return nil
}

func errwrap() {
	db := DB{driverName: "mysql", dsn: "demo:demo@tcp(127.0.0.1:3306)/oms"}
	err := db.Load()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Exec(queryNoRow)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("exec failed with no rows and %+v\n", err)
	}

	err = db.TxExec(doInsert)
	if err != nil {
		fmt.Printf("exec failed with doInsert and %+v\n", err)
	}

	err = db.TxExec(doUpdate)
	if err != nil {
		fmt.Printf("exec failed with doInsert and %+v\n", err)
	}
}
