package db

import (
	"time"

	"gorm.io/gorm"
)

type Conf struct {
	Dsn         string
	MaxIdle     int
	MaxOpen     int
	MaxLifeTime int
	MaxIdleTime int
}

func NewDb(conf *Conf) *gorm.DB {

	instance := New()

	if err := instance.Connect(conf.Dsn, &gorm.Config{}); err != nil {
		panic(err)
	}

	if err := instance.EnableConnectionPool(conf.MaxIdle, conf.MaxOpen, time.Duration(conf.MaxLifeTime)*time.Second, time.Duration(conf.MaxIdleTime)*time.Second); err != nil {
		panic(err)
	}

	return instance.DB
}
