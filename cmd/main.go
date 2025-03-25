package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/VadimGossip/extRoutingClientTester/internal/app"
	"github.com/VadimGossip/extRoutingClientTester/internal/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var appName = "Mass showroute"
var logLevel = flag.String("l", "info", "log level")

func main() {
	ctx := context.Background()
	logger.Init(getCore(getAtomicLevel(*logLevel)))

	err := initEnv()
	if err != nil {
		logger.Fatalf("failed to init env app[%s]: %s", appName, err)
	}

	a, err := app.NewApp(ctx, appName, time.Now())
	if err != nil {
		logger.Fatalf("failed to init app[%s]: %s", appName, err)
	}

	if err = a.Run(ctx); err != nil {
		logger.Fatalf("app[%s] run process finished with error: %s", appName, err)
	}
}

func initEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func getCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
	)
}

func getAtomicLevel(loglevel string) zap.AtomicLevel {
	var level zapcore.Level
	if err := level.Set(loglevel); err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}
