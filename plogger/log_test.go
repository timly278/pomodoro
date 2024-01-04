package plogger

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := New("pomodoro")

	logger.Info("tulb")
	sugar := logger.Sugar()
	sugar.Info("tick")
	logger.Info("lytu")

}
