package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/VadimGossip/extRoutingClientTester/internal/logger"
	"github.com/VadimGossip/platform_common/pkg/closer"
	"go.uber.org/zap"
)

type App struct {
	serviceProvider *serviceProvider
	name            string
	appStartedAt    time.Time
}

func NewApp(ctx context.Context, name string, appStartedAt time.Time) (*App, error) {
	a := &App{
		name:         name,
		appStartedAt: appStartedAt,
	}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	if err := a.serviceProvider.PostroutingService().RunTests(ctx); err != nil {
		logger.Error("App error",
			zap.String("method", "Run"),
			zap.String("problem", "PostroutingService RunTests"),
			zap.Error(err),
		)
	}

	gracefulShutdown(ctx, cancel, wg)
	return nil
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		logger.Info("terminating: context cancelled")
	case c := <-waitSignal():
		logger.Infof("terminating: got signal: [%s]", c)
	}

	cancel()
	if wg != nil {
		wg.Wait()
	}
}

func waitSignal() chan os.Signal {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	return sigterm
}
