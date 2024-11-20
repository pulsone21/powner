package middleware

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

type LoggerConfig struct {
	Opt    slog.HandlerOptions
	Type   string
	Writer io.Writer
}

func NewLoggerConfig(envs map[string]string) LoggerConfig {
	t := envs["LOGGER_TYPE"]
	if t == "" {
		t = "TEXT"
	}

	w_type := envs["LOGGER_OUTPUT"]
	var wr io.Writer
	switch strings.ToLower(w_type) {
	case "file":
		// TODO: Implement file logging
		wr = os.Stdout
	default:
		wr = os.Stdout
	}

	level := envs["LOGGER_LEVEL"]
	var l slog.Level
	switch strings.ToLower(level) {
	case "debug":
		l = -4
	case "error":
		l = 8
	case "warning":
		l = 4
	default:
		l = 0
	}

	source := envs["LOGGER_SOURCE"]
	s := strings.ToLower(source) == "true"

	return LoggerConfig{
		Writer: wr,
		Type:   t,
		Opt: slog.HandlerOptions{
			AddSource: s,
			Level:     l,
		},
	}
}

func Logger(config LoggerConfig) (Middleware, error) {
	if strings.ToLower(config.Type) != "json" && strings.ToLower(config.Type) != "text" {
		return nil,
			fmt.Errorf("Logger type: %v not available\n", config.Type)
	}
	return loggingMiddleware{
		opts:       config.Opt,
		writer:     config.Writer,
		loggerType: config.Type,
	}, nil
}

type loggingMiddleware struct {
	loggerType string
	writer     io.Writer
	opts       slog.HandlerOptions
}

func (l loggingMiddleware) createLogger(reqGroup slog.Attr) *slog.Logger {
	var h slog.Handler
	switch l.loggerType {
	case "JSON":
		h = slog.NewJSONHandler(l.writer, &l.opts)
	case "TEXT":
		h = slog.NewTextHandler(l.writer, &l.opts)
	}
	return slog.New(h).With(reqGroup)
}

func (l loggingMiddleware) apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqGroup := slog.Group("request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("requester", r.RemoteAddr),
			slog.String("request_id", r.Context().Value("requestID").(string)),
		)

		ctx := context.WithValue(r.Context(), "logger", l.createLogger(reqGroup))
		newReq := r.WithContext(ctx)

		log := GetLogger(newReq.Context())

		start := time.Now()
		log.Info("Incoming request")
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, newReq)

		log.Info(
			"Request handeld.",
			slog.Int("statusCode", wrapped.statusCode),
			slog.Any("duration", time.Since(start)),
		)
	})
}

func GetLogger(ctx context.Context) *slog.Logger {
	if ctx.Value("logger") != nil {
		l := ctx.Value("logger")
		return l.(*slog.Logger)
	}
	return slog.Default()
}
