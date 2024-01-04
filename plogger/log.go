package plogger

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	FILE_LOG_LEVEL     = zapcore.InfoLevel
	CONSOLSE_LOG_LEVEL = zapcore.DebugLevel
)

func New(filelogName string) *zap.Logger {
	// TODO: must understand step by step, why did you have to do these step orderly?

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	jsonEncoder := zapcore.NewJSONEncoder(encoderCfg)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, configureWriter(filelogName), FILE_LOG_LEVEL),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), CONSOLSE_LOG_LEVEL),
	)
	loggerzap := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return loggerzap
}


func configureWriter(filelogName string) zapcore.WriteSyncer {

	// bind the custom logger to zapcore
	path := filepath.Join("./logs/", filelogName+".log")
	fmt.Println("path:", path)
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    1, // megabytes
		MaxBackups: 2,
	})
}
