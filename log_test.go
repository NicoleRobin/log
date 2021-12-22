package log

import (
	"testing"
)

func TestLevel_debug(t *testing.T) {
	SetLevel(LevelDebug)
	Debug("test debug")
	Info("test debug")
	Warn("test warn")
	Error("test error")
}
