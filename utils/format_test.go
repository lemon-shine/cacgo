package utils

import (
	"testing"
)

func TestFormatFloat64ToPercentStr(t *testing.T) {
	t.Log(FormatPercentFloat64(0.12345))
}
