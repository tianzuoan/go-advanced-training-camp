package setting

import "time"

const SERVER_RUN_MODE_DEBUG = "debug"

type ServerSetting struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MysqlSetting struct {
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vip.UnmarshalKey(k, v)
}
