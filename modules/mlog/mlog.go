package mlog

type MLog struct {
	Encoder EncoderConfig
	Rotate  RotateConfig
	Level   int8
}

type EncoderConfig struct {
	TimeKey       string
	LevelKey      string
	NameKey       string
	CallerKey     string
	MessageKey    string
	StacktraceKey string
	Level         string
	Time          string
	Duration      string
	Caller        string
	Encoding      string
}

type RotateConfig struct {
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}
