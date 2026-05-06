package config

type Config struct {
	App      AppConfig   `yaml:"app"`
	Database MysqlConfig `yaml:"database"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type MysqlConfig struct {
	Driver       string `yaml:"driver"`         // 数据库驱动
	Host         string `yaml:"host"`           // 数据库地址
	Port         int    `yaml:"port"`           // 数据库端口
	Database     string `yaml:"database"`       // 数据库名称
	Username     string `yaml:"username"`       // 数据库用户名
	Password     string `yaml:"password"`       // 数据库密码
	MaxIdleConns int    `yaml:"max_idle_conns"` // 空闲连接数
	MaxOpenConns int    `yaml:"max_open_conns"` // 最大连接数
}
