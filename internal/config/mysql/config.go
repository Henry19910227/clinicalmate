package mysql

import model "clinicalmate/internal/model/config"

type config struct {
	cfg *model.MysqlConfig
}

func New(cfg *model.MysqlConfig) Config {
	return &config{cfg: cfg}
}

func (c *config) Host() string {
	return c.cfg.Host
}

func (c *config) Port() int {
	return c.cfg.Port
}

func (c *config) Database() string {
	//TODO implement me
	panic("implement me")
}

func (c *config) Username() string {
	//TODO implement me
	panic("implement me")
}

func (c *config) Password() string {
	//TODO implement me
	panic("implement me")
}

func (c *config) MaxIdleConns() int {
	//TODO implement me
	panic("implement me")
}

func (c *config) MaxOpenConns() int {
	//TODO implement me
	panic("implement me")
}
