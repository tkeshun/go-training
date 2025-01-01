package logger

import (
	"context"
	"log/slog"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// グローバル変数をリセットするためのヘルパー
var resetGlobals = func() {
	once = sync.Once{}
	globalLogger = nil
}

func TestInitLogger(t *testing.T) {
	resetGlobals()
	tests := []struct {
		name   string
		config *LoggerConfig
	}{
		{
			name:   "Normal case - valid config",
			config: &LoggerConfig{Level: slog.LevelInfo, Output: os.Stdout, CommonFields: map[string]any{"service": "test"}},
		},
		{
			name:   "Boundary case - nil config",
			config: nil, // Expected to use defaultConfig
		},
		{
			name:   "Extreme case - invalid log level",
			config: &LoggerConfig{Level: slog.Level(-100)}, // Invalid levels are ignored, treated as lowest priority
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// グローバル状態をリセット
			resetGlobals()

			// 初回呼び出し
			logger := InitLogger(tt.config)
			assert.NotNil(t, logger, "Logger should not be nil after initialization")

			// 二度目の呼び出しで再初期化されないことを確認
			secondLogger := InitLogger(tt.config)
			assert.NotNil(t, secondLogger, "Logger should still be available")
			assert.Equal(t, logger, secondLogger, "Logger instance should be the same after multiple calls")
		})
	}
}

func TestGetLogger(t *testing.T) {
	resetGlobals()
	t.Run("Error case - logger not initialized", func(t *testing.T) {
		_, err := GetLogger()
		assert.Error(t, err)
	})
	resetGlobals()
	t.Run("Normal case - logger initialized", func(t *testing.T) {
		InitLogger(nil)
		logger, err := GetLogger()
		assert.NoError(t, err)
		assert.NotNil(t, logger)
	})
}

func Test_logger_WithContext(t *testing.T) {
	type fields struct {
		contextKeys []contextKey
		cache       sync.Map
	}
	tests := []struct {
		name   string
		fields fields
		ctx    context.Context
	}{
		{
			name:   "Normal case - context with expected key",
			fields: fields{contextKeys: []contextKey{"request_id"}},
			ctx:    context.WithValue(context.Background(), contextKey("request_id"), "12345"),
		},
		{
			name:   "Boundary case - empty context",
			fields: fields{contextKeys: []contextKey{"request_id"}},
			ctx:    context.Background(),
		},
		{
			name:   "Extreme case - context with large number of keys",
			fields: fields{contextKeys: []contextKey{"key1", "key2"}},
			ctx: func() context.Context {
				ctx := context.Background()
				for i := 0; i < 1000; i++ {
					ctx = context.WithValue(ctx, contextKey("key"+strconv.Itoa(i)), i)
				}
				return ctx
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseLogger := slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
					Level: slog.LevelInfo,
				}),
			)
			logger := &Logger{
				contextKeys: tt.fields.contextKeys,
				baseLogger:  baseLogger,
			}
			cachedLogger := logger.WithContext(tt.ctx)
			assert.NotNil(t, cachedLogger)
		})
	}
}

func Test_logger_extractFieldsFromContext(t *testing.T) {
	tests := []struct {
		name           string
		contextKeys    []contextKey
		ctx            context.Context
		expectedFields map[string]any
	}{
		{
			name:           "Normal case - context with key",
			contextKeys:    []contextKey{"key1"},
			ctx:            context.WithValue(context.Background(), contextKey("key1"), "value1"),
			expectedFields: map[string]any{"key1": "value1"},
		},
		{
			name:           "Boundary case - empty context",
			contextKeys:    []contextKey{"key1"},
			ctx:            context.Background(),
			expectedFields: map[string]any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &Logger{contextKeys: tt.contextKeys}
			fields := logger.extractFieldsFromContext(tt.ctx)
			assert.Equal(t, tt.expectedFields, fields)
		})
	}
}

func Test_logger_log(t *testing.T) {
	t.Run("Normal case - log info level", func(t *testing.T) {
		logger := &Logger{baseLogger: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
		assert.NotPanics(t, func() {
			logger.log(context.TODO(), slog.LevelInfo, "Test message", map[string]any{"key": "value"})
		})
	})
}
