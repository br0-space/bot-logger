package logger_test

import (
	"reflect"
	"testing"

	logger "github.com/br0-space/bot-logger"
	"github.com/br0-space/bot-logger/null"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want logger.Interface
	}{
		{"null", null.New()},
	}
	for _, tt := range tests {
		it := tt
		t.Run(it.name, func(t *testing.T) {
			t.Parallel()
			if got := logger.New(); !reflect.DeepEqual(got, it.want) {
				t.Errorf("New() = %v, want %v", got, it.want)
			}
		})
	}
}
