package package_core

import (
	"fmt"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Core struct {
	Config *Config
	Engine *xorm.Engine
}

func NewCore() (*Core, error) {
	conf := NewConfig()

	engineSource := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		conf.DB.Driver,
		conf.DB.User,
		conf.DB.Password,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.Database)
	engine, err := xorm.NewEngine(conf.DB.Driver, engineSource)

	if err != nil {
		return nil, err
	}

	return &Core{
		Config: conf,
		Engine: engine,
	}, nil
}
