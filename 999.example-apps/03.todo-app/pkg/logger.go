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

// グローバル変数はmain関数実行前に宣言される。
var globalLogger *logger

// sync.Onceオブジェクトに対して渡した関数は、プログラム全体で一度しか実行されない
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

// ロガーを初期化する関数
func InitLogger(config *LoggerConfig) *logger {
	// もう一度onceが実行されるとエラーになる
	once.Do(
		func() {
			// 初期設定が与えられていない場合、初期設定を使う。
			if config == nil {
				config = &defaultConfig
			}
			// ロガーを新規に作成
			baseLogger := slog.New(
				slog.NewJSONHandler(config.Output, &slog.HandlerOptions{
					Level: config.Level,
				}),
			)
			// 共通フィールドを追加する
			for k, v := range config.CommonFields {
				baseLogger = baseLogger.With(slog.Any(k, v))
			}
			// グローバル変数にあるlogger格納用変数にインスタンスを格納
			globalLogger = &logger{
				baseLogger:  baseLogger,
				contextKeys: config.ContextKeys,
			}
		},
	)
	return globalLogger
}

func GetLogger() (*logger, error) {
	// もし初期化されてなかったらエラー
	if globalLogger == nil {
		return nil, fmt.Errorf("error: no init logger. Please initialize Logger. Using InitLogger")
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

// ログを処理
func (l *logger) log(ctx context.Context, level slog.Level, msg string, fields map[string]any) {
	args := make([]any, 0, len(fields)*2)
	for k, v := range fields {
		args = append(args, k, v)
	}
	l.baseLogger.Log(ctx, level, msg, args...)
}

// ログ出力メソッド群
func (l *logger) Info(ctx context.Context, msg string, fields map[string]any) {
	l.log(ctx, slog.LevelInfo, msg, fields)
}

func (l *logger) Warn(ctx context.Context, msg string, fields map[string]any) {
	l.log(ctx, slog.LevelWarn, msg, fields)
}

func (l *logger) Error(ctx context.Context, msg string, fields map[string]any) {
	l.log(ctx, slog.LevelError, msg, fields)
}

func (l *logger) Debug(ctx context.Context, msg string, fields map[string]any) {
	l.log(ctx, slog.LevelDebug, msg, fields)
}
