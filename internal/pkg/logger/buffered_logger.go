package logger

import (
	"bytes"
	goLogger "log"
	"os"
)

type BufferLogger struct {
	loggerBuff *goLogger.Logger
	logger     *goLogger.Logger
	Out        *bytes.Buffer
	printToIO  bool
	level      Level
}

func NewBufferLogger(printToIO bool) ILogger {
	outBuff := new(bytes.Buffer)
	stdLogg := goLogger.New(os.Stdout, "", 0)

	return &BufferLogger{
		loggerBuff: goLogger.New(outBuff, "", 0),
		logger:     stdLogg,
		Out:        outBuff,
		printToIO:  printToIO,
	}
}

func (bl *BufferLogger) Trace(args ...interface{}) {
	bl.loggerBuff.Println("TRACE ", args)
	if bl.printToIO {
		bl.logger.Println("TRACE ", args)
	}
}

func (bl *BufferLogger) Debug(args ...interface{}) {
	bl.loggerBuff.Println("DEBUG ", args)
	if bl.printToIO {
		bl.logger.Println("DEBUG ", args)
	}
}

func (bl *BufferLogger) Info(args ...interface{}) {
	bl.loggerBuff.Println("INFO ", args)
	if bl.printToIO {
		bl.logger.Println("INFO ", args)
	}
}

func (bl *BufferLogger) Warn(args ...interface{}) {
	bl.loggerBuff.Println("WARN ", args)
	if bl.printToIO {
		bl.logger.Println("WARN ", args)
	}
}

func (bl *BufferLogger) Error(args ...interface{}) {
	bl.loggerBuff.Println("ERROR ", args)
	if bl.printToIO {
		bl.logger.Println("ERROR ", args)
	}
}

func (bl *BufferLogger) Panic(args ...interface{}) {
	bl.loggerBuff.Println("PANIC ", args)
	if bl.printToIO {
		bl.logger.Println("PANIC ", args)
	}
}

func (bl *BufferLogger) Fatal(args ...interface{}) {
	bl.loggerBuff.Println("FATAL ", args)
	if bl.printToIO {
		bl.logger.Println("FATAL ", args)
	}
}

func (bl *BufferLogger) GetBufferedString() string {
	return bl.Out.String()
}

func (bl *BufferLogger) TraceIsEnabled() bool {
	return bl.level.Enabled(TraceLevel)
}

func (bl *BufferLogger) DebugIsEnabled() bool {
	return bl.level.Enabled(DebugLevel)
}

func (bl *BufferLogger) InfoIsEnabled() bool {
	return bl.level.Enabled(InfoLevel)
}

func (bl *BufferLogger) WarnIsEnabled() bool {
	return bl.level.Enabled(WarnLevel)
}

func (bl *BufferLogger) ErrorIsEnabled() bool {
	return bl.level.Enabled(ErrorLevel)
}

func (bl *BufferLogger) PanicIsEnabled() bool {
	return bl.level.Enabled(PanicLevel)
}

func (bl *BufferLogger) FatalIsEnabled() bool {
	return bl.level.Enabled(FatalLevel)
}
