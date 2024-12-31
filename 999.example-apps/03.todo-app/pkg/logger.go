package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type contextKey string

// json形式での構造化ログを出力するロガー
type LoggerConfig struct {
	Level        slog.Level     // ログレベル
	CommonFields map[string]any // 共通フィールド
	Output       *os.File       // 出力先
	ContextKeys  []contextKey   //コンテキストから取得するキー一覧
}

type logger struct {
	baseLogger  *slog.Logger
	contextKeys []contextKey
	cache       sync.Map // キャッシュ
}

var globalLogger *logger
var once sync.Once

var defaultConfig = LoggerConfig{
	Level: slog.LevelInfo,
	CommonFields: map[string]any{
		"service": "no setting",
		"env":     "no setting",
	},
	Output:      os.Stdout,
	ContextKeys: []contextKey{"request_id"},
}

func InitLogger(config *LoggerConfig) (*logger, error) {
	once.Do(
		func() {
			if config == nil {
				config = &defaultConfig
			}

			baseLogger := slog.New(
				slog.NewJSONHandler(config.Output, &slog.HandlerOptions{
					Level: config.Level,
				}),
			)
			// 共通フィールドを追加する
			for k, v := range config.CommonFields {
				baseLogger = baseLogger.With(slog.Any(k, v))
			}

			globalLogger = &logger{
				baseLogger:  baseLogger,
				contextKeys: config.ContextKeys,
			}
		},
	)
	return globalLogger, nil
}

func GetLogger() (*logger, error) {
	if globalLogger == nil {
		return nil, fmt.Errorf("error: no init logger. Please initialize Logger. Using InitLogger.")
	}
	return globalLogger, nil
}

// コンテキスト情報を出力できるロガーを返す。
func (l *logger) WithContext(ctx context.Context) *logger {
	//  キャッシュをチェックする
	// 同じコンテキストでのロガー呼び出しではキャッシュが使われる。
	if cached, ok := l.cache.Load(ctx); ok {
		return cached.(*logger)
	}

	// contextにあるvalueを取得
	fields := l.extractFieldsFromContext(ctx)
	// 現在のcontextからロガーを切り離す。
	newLogger := l.baseLogger
	// 取得したキーとバリューをslogの出力に追加
	for k, v := range fields {
		newLogger = newLogger.With(slog.Any(k, v))
	}
	// 新しいcontextで使うロガーを生成
	cachedLogger := &logger{
		baseLogger:  newLogger,     // 新しく作ったロガーを渡す。
		contextKeys: l.contextKeys, // 初期設定のキーバリューを渡す。
	}

	l.cache.Store(ctx, cachedLogger) // ここまで到達したということはキャッシュがないので格納する。
	return cachedLogger
}

// 事前定義したキーに対応する値をコンテキストから取得する
func (l *logger) extractFieldsFromContext(ctx context.Context) map[string]any {
	fields := make(map[string]any, len(l.contextKeys))
	for _, key := range l.contextKeys {
		if value := ctx.Value(key); value != nil {
			fields[string(key)] = value
		}
	}
	return fields
}

// log は内部でログを処理します。
func (l *logger) log(level slog.Level, msg string, fields map[string]any) {
	args := make([]any, 0, len(fields)*2)
	for k, v := range fields {
		args = append(args, k, v)
	}
	l.baseLogger.Log(level, msg, args...)
}

// ログ出力メソッド群
func (l *logger) Info(msg string, fields map[string]any) {
	l.log(slog.LevelInfo, msg, fields)
}

func (l *logger) Warn(msg string, fields map[string]any) {
	l.log(slog.LevelWarn, msg, fields)
}

func (l *logger) Error(msg string, fields map[string]any) {
	l.log(slog.LevelError, msg, fields)
}

func (l *logger) Debug(msg string, fields map[string]any) {
	l.log(slog.LevelDebug, msg, fields)
}
