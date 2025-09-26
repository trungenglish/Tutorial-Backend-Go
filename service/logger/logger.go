package logger

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

var Log zerolog.Logger

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "message"

	Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func LogWithTrace(ctx context.Context) *zerolog.Logger {
	span := trace.SpanFromContext(ctx)
	if !span.SpanContext().IsValid() {
		return &Log
	}

	spanCtx := span.SpanContext()
	logger := Log.With().
		Str("trace_id", spanCtx.TraceID().String()).
		Str("span_id", spanCtx.SpanID().String()).
		Logger()

	return &logger
}

func Info(ctx context.Context) *zerolog.Event {
	return LogWithTrace(ctx).Info()
}

func Error(ctx context.Context) *zerolog.Event {
	return LogWithTrace(ctx).Error()
}

func Warn(ctx context.Context) *zerolog.Event {
	return LogWithTrace(ctx).Warn()
}

func Debug(ctx context.Context) *zerolog.Event {
	return LogWithTrace(ctx).Debug()
}
