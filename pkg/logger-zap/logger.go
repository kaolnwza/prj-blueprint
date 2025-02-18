package logger

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kaolnwza/proj-blueprint/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerKey struct{}

const (
	levelDevelopment = "development"
	levelProduction  = "production"
)

const (
	envLocal = "local"
)

func defaultConfig(config zap.Config) zap.Config {
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.NameKey = "id"
	config.DisableStacktrace = true
	return config
}

func InitZapLogger(conf config.LogConfig, env string) {
	config := zap.NewProductionConfig()
	config.EncoderConfig = newZapEncoderConfig()
	setLevel(&config, conf.Level)
	setFormat(&config, conf.JsonFormat)

	if env == envLocal {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	l, _ := defaultConfig(config).Build()
	zap.ReplaceGlobals(l)
}

func setLevel(conf *zap.Config, lv string) {
	logLv, err := zap.ParseAtomicLevel(lv)
	if err != nil {
		log.Fatalf("fail to init log: %v", err)
	}

	conf.Level.SetLevel(logLv.Level())
}

func setFormat(conf *zap.Config, isJsonFormat bool) {
	if isJsonFormat {
		conf.Encoding = "json"
	}
}

func GetByContext(ctx context.Context) *zap.SugaredLogger {
	if id, ok := ctx.Value(loggerKey{}).(string); ok {
		return zap.L().Sugar().Named(id)
	}

	return zap.L().Sugar()
}

// init to setup log id per request-id or kafka-message-id
func New(ctx *context.Context) *zap.SugaredLogger {
	ctxT := *ctx
	id, ok := ctxT.Value(loggerKey{}).(string)
	if !ok {
		id = uuid.New().String()
		ctxT = context.WithValue(ctxT, loggerKey{}, id)
		*ctx = ctxT
	}

	return zap.L().Sugar().Named(id)
}
