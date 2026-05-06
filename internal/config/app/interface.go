package app

type Config interface {
	Name() string
	Ip() string
	Port() int
}
