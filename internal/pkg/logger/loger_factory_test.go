package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tlLogger = LoggerFactory(WithBufferedLoggerImplementer(), WithPrintToIO())

func TestLoggerFactory(t *testing.T) {
	type args struct {
		option []Option
	}
	tests := []struct {
		name        string
		args        args
		wantResults map[string]string
	}{
		{
			name: "TestLoggerFactory logger",
			args: args{option: []Option{
				WithBufferedLoggerImplementer(),
				WithLoggerLevel("TRACE"),
			},
			},
			wantResults: map[string]string{"TRACE": "trace msg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := LoggerFactory(tt.args.option...)

			if logger.TraceIsEnabled() {
				logger.Trace(tt.wantResults["TRACE"])
				assert.Equal(t, tt.wantResults["TRACE"], (logger).(*BufferLogger).GetBufferedString())
			}

		})
	}
}
