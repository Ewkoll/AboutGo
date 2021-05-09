package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/ewkoll/aboutgo/server"
	"golang.org/x/sync/errgroup"
)

type App struct {
	servers []server.Server
	ctx     context.Context
	sigs    []os.Signal
	cancel  func()
}

func New() *App {
	ctx_b := context.Background()
	sigs := []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
	ctx, cancel := context.WithCancel(ctx_b)
	app := &App{
		ctx:     ctx,
		sigs:    sigs,
		cancel:  cancel,
		servers: []server.Server{server.CreateHttpServer(":9090", ctx), server.CreateTcpServer(":9091", ctx)},
	}
	return app
}

func (a *App) Run() error {
	egroup, ctx := errgroup.WithContext(a.ctx)
	for _, srv := range a.servers {
		srv := srv
		egroup.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
		egroup.Go(func() error {
			return srv.Start()
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.sigs...)
	egroup.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.Stop()
			}
		}
	})
	if err := egroup.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
