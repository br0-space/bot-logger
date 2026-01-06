package null

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(_ ...any) {}

func (l *Logger) Debugf(_ string, _ ...any) {}

func (l *Logger) Info(_ ...any) {}

func (l *Logger) Infof(_ string, _ ...any) {}

func (l *Logger) Warning(_ ...any) {}

func (l *Logger) Warningf(_ string, _ ...any) {}

func (l *Logger) Error(_ ...any) {}

func (l *Logger) Errorf(_ string, _ ...any) {}

func (l *Logger) Panic(_ ...any) {}

func (l *Logger) Panicf(_ string, _ ...any) {}

func (l *Logger) Fatal(_ ...any) {}

func (l *Logger) Fatalf(_ string, _ ...any) {}

func (l *Logger) SetExtraCallDepth(_ int) {}

func (l *Logger) ResetExtraCallDepth() {}
