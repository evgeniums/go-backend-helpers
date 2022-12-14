package logger

type Fields = map[string]interface{}
type Level int

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

type Logger interface {
	Config() interface{}

	Log(level Level, message string, fields ...Fields)

	Error(message string, err error, fields ...Fields) error
	Warn(message string, fields ...Fields)
	Debug(message string, fields ...Fields)
	Info(message string, fields ...Fields)
	Fatal(message string, err error, fields ...Fields) error

	ErrorRaw(...interface{})
}

type WithLogger interface {
	Logger() Logger
}

type WithLoggerBase struct {
	logger Logger
}

func (w *WithLoggerBase) Logger() Logger {
	return w.logger
}

func (w *WithLoggerBase) SetLogger(logger Logger) {
	w.logger = logger
}
