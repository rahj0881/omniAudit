package log

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync"

	"github.com/omni-network/omni/lib/errors"

	charm "github.com/charmbracelet/log"
	"github.com/muesli/termenv"
)

//nolint:gochecknoglobals // Global logger is our approach.
var (
	global   = newConsoleLogger()
	globalMu = new(sync.RWMutex)
)

type loggerKey struct{}

// Init initializes the global logger with the given config.
// It also returns a copy of the context with the global attached, see WithLogger.
// It returns an error if the config is invalid.
func Init(ctx context.Context, cfg Config) (context.Context, error) {
	l, err := cfg.make()
	if err != nil {
		return nil, err
	}

	globalMu.Lock()
	global = l
	globalMu.Unlock()

	return WithLogger(ctx, l), nil
}

// WithLogger returns a copy of the context with which the global
// is associated replacing the default global logger when logging with this context.
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// GetLogger returns the logger from the context, or the global logger if not present.
func GetLogger(ctx context.Context) *slog.Logger {
	if l := ctx.Value(loggerKey{}); l != nil {
		return l.(*slog.Logger) //nolint:forcetypeassert,revive // We know the type.
	}

	globalMu.RLock()
	defer globalMu.RUnlock()

	return global
}

func newJSONLogger(opts ...func(*options)) *slog.Logger {
	o := defaultOptions()
	for _, opt := range opts {
		opt(&o)
	}

	// Maybe replace time, source and stacktrace with stubs for testing
	replaceAttr := func(groups []string, a slog.Attr) slog.Attr { return a }
	if o.Test {
		replaceAttr = func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.String(slog.TimeKey, "00-00-00 00:00:00")
			}
			if a.Key == slog.SourceKey && len(groups) == 0 {
				return slog.String(slog.SourceKey, "<source>")
			}
			if a.Key == "stacktrace" && len(groups) == 0 {
				return slog.String(slog.SourceKey, "<stacktrace>")
			}

			return a
		}
	}

	handler := slog.NewJSONHandler(o.Writer, &slog.HandlerOptions{
		AddSource:   true,
		Level:       o.Level,
		ReplaceAttr: replaceAttr,
	})

	return slog.New(handler)
}

// newConsoleLogger returns a new console global for the following opinionated style:
// - Colored log levels (if tty supports it)
// - Timestamps are concise with millisecond precision
// - Timestamps and structured keys are faint
// - Messages are right padded to 40 characters
// This is aimed at local-dev and debugging. Production should use json or logfmt.
func newConsoleLogger(opts ...func(*options)) *slog.Logger {
	o := defaultOptions()
	for _, opt := range opts {
		opt(&o)
	}

	timeFormat := "06-01-02 15:04:05.000"
	if o.Test {
		timeFormat = "00-00-00 00:00:00"
	}

	charmLevel, _ := charm.ParseLevel(o.Level.String()) // Ignore error as all slog levels are valid charm levels.

	logger := charm.NewWithOptions(o.Writer, charm.Options{
		TimeFormat:      timeFormat,
		ReportTimestamp: true,
		Level:           charmLevel,
	})

	styles := charm.DefaultStyles()
	styles.Timestamp = styles.Timestamp.Faint(true)
	const padWidth = 40
	styles.Message = styles.Message.Width(padWidth).Inline(true)
	logger.SetStyles(styles)
	logger.SetColorProfile(o.Color)

	if o.Test {
		return slog.New(stubHandler{Handler: logger})
	}

	return slog.New(logger)
}

// options configure new loggers.
type options struct {
	Writer io.Writer // Write to some buffer
	Level  slog.Level
	Color  termenv.Profile
	Test   bool // Stubs non-deterministic output for tests.
}

func defaultOptions() options {
	return options{
		Writer: os.Stderr,
		Level:  slog.LevelDebug,
		Color:  termenv.ColorProfile(),
		Test:   false,
	}
}

// WithNoopLogger returns a copy of the context with a noop global which discards all logs.
func WithNoopLogger(ctx context.Context) context.Context {
	return WithLogger(ctx, newConsoleLogger(func(o *options) {
		o.Writer = io.Discard
	}))
}

// stubHandler is a handler that replaces the stacktrace and source attributes with stubs.
type stubHandler struct {
	slog.Handler
}

func (t stubHandler) Handle(ctx context.Context, r slog.Record) error {
	resp := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)

	r.Attrs(func(a slog.Attr) bool {
		if a.Key == "stacktrace" {
			resp.AddAttrs(slog.String("stacktrace", "<stacktrace>"))
		} else {
			resp.AddAttrs(a)
		}

		return true
	})

	if err := t.Handler.Handle(ctx, resp); err != nil {
		return errors.Wrap(err, "handle")
	}

	return nil
}
