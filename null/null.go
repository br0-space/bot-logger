package null

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(args ...interface{}) {}

func (l *Logger) Debugf(format string, args ...interface{}) {}

func (l *Logger) Info(args ...interface{}) {}

func (l *Logger) Infof(format string, args ...interface{}) {}

func (l *Logger) Warning(args ...interface{}) {}

func (l *Logger) Warningf(format string, args ...interface{}) {}

func (l *Logger) Error(args ...interface{}) {}

func (l *Logger) Errorf(format string, args ...interface{}) {}

func (l *Logger) Panic(args ...interface{}) {}

func (l *Logger) Panicf(format string, args ...interface{}) {}

func (l *Logger) Fatal(args ...interface{}) {}

func (l *Logger) Fatalf(format string, args ...interface{}) {}

func (l *Logger) SetExtraCallDepth(depth int) {}

func (l *Logger) ResetExtraCallDepth() {}
