package loggerx

import (
	"log/slog"
	"os"

	"github.com/Rephobia/green-api-test-task/internal/config"
)

func New(config *config.Config) *slog.Logger {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:       slog.LevelInfo,
			AddSource:   true,
			ReplaceAttr: nil,
		}),
	).With(
		AppNameField(config.AppName),
		EnvField(config.Env),
	)

	slog.SetDefault(logger)

	return logger
}

func AppNameField(value string) slog.Attr { return slog.String("app_name", value) }
func AddrField(value string) slog.Attr    { return slog.String("addr", value) }
func EnvField(value string) slog.Attr     { return slog.String("env", value) }
func ErrorField(value error) slog.Attr    { return slog.Any("error", value) }
