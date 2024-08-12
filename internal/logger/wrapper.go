package logger

import (
	"context"
	"log/slog"
	"os"
)

// Logger - wrapper для библиотеки логирования
type Logger interface {
	Warning(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Info(...interface{})
	Debug(...interface{})
	Fatal(...interface{})
	GetLogger() *slog.Logger
}

type wrapper struct {
	ctx context.Context
	log *slog.Logger
}

// New
func New(ctx context.Context, log *slog.Logger) Logger {
	return &wrapper{
		ctx: ctx,
		log: log,
	}
}

func (log *wrapper) Warning(args ...interface{}) {
	log.Warning(log.ctx, args)
	log.ctx = nil
}

func (log *wrapper) Warn(args ...interface{}) {
	log.Warning(log.ctx, args)
	log.ctx = nil
}

func (log *wrapper) Error(args ...interface{}) {
	log.Error(log.ctx, args)
	log.ctx = nil
}

func (log *wrapper) Info(args ...interface{}) {
	log.Info(log.ctx, args)
	log.ctx = nil
}

func (log *wrapper) Debug(args ...interface{}) {
	log.Debug(log.ctx, args)
	log.ctx = nil
}

func (log *wrapper) Fatal(args ...interface{}) {
	log.Error(log.ctx, args)
	log.ctx = nil
	os.Exit(-1)
}

func (log *wrapper) GetLogger() *slog.Logger {
	return log.log
}
