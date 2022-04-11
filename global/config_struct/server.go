package config_struct

/*
   功能说明: server.yaml配置文件对应的结构体
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/18 10:54
*/

type ServerConfig struct {
	Name          string  `mapstructure:"name"`
	Env           string  `mapstructure:"env"`
	RunMode       string  `mapstructure:"runMode"`
	SystemVersion string  `mapstructure:"version"`
	MySQL         Mysql   `mapstructure:"mysql"`
	Redis         Redis   `mapstructure:"redis"`
	Service       Service `mapstructure:"service"`
	Log           Log
}

type Mysql struct {
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Ip        string `mapstructure:"ip"`
	Port      int    `mapstructure:"port"`
	Db        string `mapstructure:"db"`
	Parameter string `mapstructure:"parameter"`
}

type Redis struct {
}

type Service struct {
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
}

type Log struct {
	LogDir     string // 日志保存的目录
	InfoLog    string // 正常日志文件名
	ErrorLog   string // 错误日志文件名
	MaxSize    int    // 在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups int    // 保留旧文件的最大个数
	MaxAge     int    // 保留旧文件的最大天数
	Compress   bool   // 是否压缩/归档旧文件
}
