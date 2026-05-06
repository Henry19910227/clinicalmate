package mysql

type Config interface {
	Host() string
	Port() int
	Database() string
	Username() string
	Password() string
	MaxIdleConns() int
	MaxOpenConns() int
}
