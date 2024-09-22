package utils

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func InitTintedSlog(verbose bool) {
	noColor := false
	noColorEnv, _ := os.LookupEnv("NOCOLOR")
	if noColorEnv == "1" {
		noColor = true
	}

	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}

	pretty := tint.NewHandler(os.Stderr, &tint.Options{
		Level:      level,
		TimeFormat: time.Kitchen,
		NoColor:    noColor,
	})
	slog.SetDefault(slog.New(pretty))
}

func Fatal(err error) {
	slog.Error("fatal error", "err", err)
	os.Exit(1)
}
