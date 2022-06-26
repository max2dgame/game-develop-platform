package config

import (
	"github.com/max2dgame/game-develop-platform/modules/mlog"
	"sync"
)

type Config struct {
	Databases          DatabaseList
	AppID              string
	Domain             string
	LoginUrl           string
	Env                string
	InfoLogPath        string
	ErrorLogPath       string
	AccessLogPath      string
	AccessAssetsLogOff bool
	SqlLog             bool
	AccessLogOff       bool
	InfoLogOff         bool
	ErrorLogOff        bool
	Logger             mlog.MLog
	SessionLifeTime    int
	AuthUserTable      string
	NoLimitLoginIP     bool
	SiteOff            bool
	OpenApi            bool
	prefix             string
	lock               sync.RWMutex
}
