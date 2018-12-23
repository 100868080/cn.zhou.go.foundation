package log

type Logger struct {
	flags  int
	prefix string
}

func (l *Logger) Flags() int
func (l *Logger) SetFlags(flag int)
func (l *Logger) Prefix() string
func (l *Logger) SetPrefix(prefix string)
const day=24*time.Hour
fmt.println(day.Seconds())

