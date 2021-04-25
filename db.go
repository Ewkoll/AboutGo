package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
)

type Callback func(db *sql.DB, ctx context.Context) error

type DB struct {
	driverName string
	dsn        string
	pool       *sql.DB
	ctx        context.Context
	cancel     context.CancelFunc
}

func (db *DB) Load() error {
	if db.driverName == "" {
		db.driverName = os.Getenv("DriverName")
	}
	if db.dsn == "" {
		db.dsn = os.Getenv("DSN")
	}

	pool, err := sql.Open(db.driverName, db.dsn)
	if err != nil {
		return xerrors.Wrap(err, "db: sql open failed")
	}

	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(10)
	pool.SetConnMaxLifetime(time.Hour * 4)
	pool.SetConnMaxIdleTime(time.Minute * 15)

	ctx, stop := context.WithCancel(context.Background())
	appSignal := make(chan os.Signal)
	signal.Notify(appSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range appSignal {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				stop()
			default:
				fmt.Printf("")
			}
		}

	}()

	db.ctx = ctx
	db.pool = pool
	db.cancel = stop
	return nil
}

func (db *DB) Ping() error {
	if db.ctx == nil || db.pool == nil {
		return xerrors.New("db: could not execute load or load failed.")
	}

	ctx, cancel := context.WithTimeout(db.ctx, 30*time.Second)
	defer cancel()

	if err := db.pool.PingContext(ctx); err != nil {
		return xerrors.Wrap(err, "db: unable to connect to database.")
	}
	return nil
}

func (db *DB) Close() {
	if db.pool == nil || db.cancel == nil {
		return
	}

	db.cancel()
	db.pool.Close()
}

func (db *DB) Exec(callback Callback) error {
	ctx, cancel := context.WithTimeout(db.ctx, 30*time.Second)
	defer cancel()
	return callback(db.pool, ctx)
}

func (db *DB) TxExec(callback Callback) error {
	ctx, cancel := context.WithTimeout(db.ctx, 30*time.Second)
	defer cancel()

	tx, err := db.pool.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return xerrors.Wrap(err, "db: does not support a given isolation level")
	}

	err = callback(db.pool, ctx)
	if err != nil {
		tx.Rollback()
		return xerrors.Wrap(err, "db: finish execute rollback")
	}

	err = tx.Commit()
	if err != nil {
		return xerrors.Wrap(err, "db: commit failed")
	}
	return nil
}
