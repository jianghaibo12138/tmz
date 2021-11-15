package statistic

import (
	"testing"
)

func Test_logPosParser(t *testing.T) {
	tests := []struct {
		name string
		want *LogPos
	}{
		{name: "c1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := logPosParser()
			t.Logf("%+v", got)
		})
	}
}
