package plogger

import (
	"pomodoro/util"
	"testing"
)

func TestLog(t *testing.T) {
	conf := &util.Config{LogFilesPath: "./logs/"}
	logger := New(conf, "pomodoro")

	logger.Info("tulb")
	sugar := logger.Sugar()
	sugar.Info("tick")
	logger.Info("lytu")

}
