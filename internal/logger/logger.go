package logger

import (
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

var log = zerolog.New(io.Discard) // default: disabled

func Init(level string) {
	lvl, err := zerolog.ParseLevel(strings.ToLower(level))

	if err != nil {
		lvl = zerolog.InfoLevel
	}

	log = zerolog.New(os.Stderr).Level(lvl).With().Timestamp().Logger()
}

func Debug() *zerolog.Event { return log.Debug() }
func Info() *zerolog.Event  { return log.Info() }
func Warn() *zerolog.Event  { return log.Warn() }
func Error() *zerolog.Event { return log.Error() }
func Fatal() *zerolog.Event { return log.Fatal() }
func Panic() *zerolog.Event { return log.Panic() }
