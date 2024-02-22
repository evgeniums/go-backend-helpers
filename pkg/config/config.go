package config

import "time"

type Config interface {
	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetUint(key string) uint
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetFloat64(key string) float64
	GetIntSlice(key string) []int
	GetStringSlice(key string) []string
	GetFloat64Slice(key string) []float64
	GetStringMapString(key string) map[string]string

	GetTime(key string) time.Time
	GetDuration(key string) time.Duration

	SetDefault(key string, value interface{})
	Set(key string, value interface{})

	IsSet(key string) bool

	AllKeys() []string

	Rebuild() error
	ToString() string
}

type WithCfg interface {
	Cfg() Config
}

type WithCfgBase struct {
	cfg Config
}

func (w *WithCfgBase) Cfg() Config {
	return w.cfg
}

func (w *WithCfgBase) SetCfg(cfg Config) {
	w.cfg = cfg
}
