package logger

import (
	"context"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/suctl/aws-powertools-lambda-go/internal/utils"
	"github.com/suctl/aws-powertools-lambda-go/logger/types"
)

const (
	callerName           = "location"
	callerSkipFrameCount = 3
	defaultLogLevel      = "DEBUG"
)

var LogMapper = map[string]zerolog.Level{
	"ERROR": zerolog.ErrorLevel,
	"WARN":  zerolog.WarnLevel,
	"INFO":  zerolog.InfoLevel,
	"DEBUG": zerolog.DebugLevel,
	"TRACE": zerolog.TraceLevel,
}

type Logger struct {
	logger zerolog.Logger
}

func (log *Logger) InjectContext(ctx context.Context) {
	lambdaContext := newLambdaContext(ctx)
	if lambdaContext.FunctionRequestId == "unknown" {
		log.Warn("failed to load function request id")
	}
	if lambdaContext.FunctionARN == "" {
		log.Warn("failed to load function arn")
	}
	if lambdaContext.FunctionName == "unknown" {
		log.Warn("failed to load function name")
	}
	if lambdaContext.FunctionMemorySize == "0" {
		log.Warn("failed to load function memory size")
	}
	log.logger = log.logger.With().
		Str("function_name", lambdaContext.FunctionName).
		Str("function_memory_size", lambdaContext.FunctionMemorySize).
		Str("function_arn", lambdaContext.FunctionARN).
		Str("function_request_id", lambdaContext.FunctionRequestId).
		Logger()
}

func (log *Logger) Error(message string, args ...any) {
	log.logger.Error().Msgf(message, args...)
}

func (log *Logger) Warn(message string, args ...any) {
	log.logger.Warn().Msgf(message, args...)
}

func (log *Logger) Info(message string, args ...any) {
	log.logger.Info().Msgf(message, args...)
}

func (log *Logger) Debug(message string, args ...any) {
	log.logger.Debug().Msgf(message, args...)
}

func (log *Logger) Trace(message string, args ...any) {
	log.logger.Trace().Msgf(message, args...)
}

func New(logConfig types.LogConfig) *Logger {
	setConfigFromEnvironment()
	config := newConfig(&logConfig)
	zerolog.CallerFieldName = callerName
	logger := zerolog.
		New(config.Writer).
		With().
		CallerWithSkipFrameCount(callerSkipFrameCount).
		Timestamp().
		Logger()

	for key, value := range logConfig.Properties {
		logger = logger.With().Str(key, value).Logger()
	}

	return &Logger{
		logger: logger,
	}
}

func setConfigFromEnvironment() {
	logLevel := utils.GetEnvironmentVariable("POWERTOOLS_LOG_LEVEL", defaultLogLevel)
	zerolog.SetGlobalLevel(LogMapper[strings.ToUpper(logLevel)])
}

func newLambdaContext(ctx context.Context) types.LambdaContext {
	lambdaContext := types.LambdaContext{
		FunctionName:       utils.GetEnvironmentVariable("AWS_LAMBDA_FUNCTION_NAME", "unknown"),
		FunctionMemorySize: utils.GetEnvironmentVariable("AWS_LAMBDA_FUNCTION_MEMORY_SIZE", "0"),
		FunctionARN:        utils.GetEnvironmentVariable("AWS_LAMBDA_FUNCTION_INVOKED_ARN", ""),
		FunctionRequestId:  "unknown",
	}
	functionRequestId := ctx.Value("function_request_id")
	if functionRequestId != nil {
		lambdaContext.FunctionRequestId = functionRequestId.(string)
	}
	return lambdaContext
}

func newConfig(logConfig *types.LogConfig) *types.LogConfig {
	if logConfig.Writer == nil {
		logConfig.Writer = os.Stdout
	}
	return logConfig
}
