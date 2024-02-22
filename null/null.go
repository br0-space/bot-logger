package null

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(_ ...interface{}) {}

func (l *Logger) Debugf(_ string, _ ...interface{}) {}

func (l *Logger) Info(_ ...interface{}) {}

func (l *Logger) Infof(_ string, _ ...interface{}) {}

func (l *Logger) Warning(_ ...interface{}) {}

func (l *Logger) Warningf(_ string, _ ...interface{}) {}

func (l *Logger) Error(_ ...interface{}) {}

func (l *Logger) Errorf(_ string, _ ...interface{}) {}

func (l *Logger) Panic(_ ...interface{}) {}

func (l *Logger) Panicf(_ string, _ ...interface{}) {}

func (l *Logger) Fatal(_ ...interface{}) {}

func (l *Logger) Fatalf(_ string, _ ...interface{}) {}

func (l *Logger) SetExtraCallDepth(_ int) {}

func (l *Logger) ResetExtraCallDepth() {}
