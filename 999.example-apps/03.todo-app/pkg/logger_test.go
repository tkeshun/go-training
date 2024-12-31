package pkg

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name    string
		config  *LoggerConfig
		wantErr bool
	}{
		{
			name:    "Normal case - valid config",
			config:  &LoggerConfig{Level: slog.LevelInfo, Output: os.Stdout, CommonFields: map[string]any{"service": "test"}},
			wantErr: false,
		},
		{
			name:    "Boundary case - nil config",
			config:  nil, // Expected to use defaultConfig
			wantErr: false,
		},
		{
			name:    "Extreme case - invalid log level",
			config:  &LoggerConfig{Level: slog.Level(-100)},
			wantErr: false, // Invalid levels are ignored, treated as lowest priority
		},
		// No abnormal case: InitLogger initializes logger with any given config
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := InitLogger(tt.config)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, logger)
			}
		})
	}
}

func TestGetLogger(t *testing.T) {
	t.Run("Error case - logger not initialized", func(t *testing.T) {
		_, err := GetLogger()
		assert.Error(t, err)
	})

	t.Run("Normal case - logger initialized", func(t *testing.T) {
		_, err := InitLogger(nil)
		assert.NoError(t, err)

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
					ctx = context.WithValue(ctx, contextKey("key"+string(i)), i)
				}
				return ctx
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &logger{contextKeys: tt.fields.contextKeys}
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
			logger := &logger{contextKeys: tt.contextKeys}
			fields := logger.extractFieldsFromContext(tt.ctx)
			assert.Equal(t, tt.expectedFields, fields)
		})
	}
}

func Test_logger_log(t *testing.T) {
	t.Run("Normal case - log info level", func(t *testing.T) {
		logger := &logger{baseLogger: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
		assert.NotPanics(t, func() {
			logger.log(slog.LevelInfo, "Test message", map[string]any{"key": "value"})
		})
	})
}
