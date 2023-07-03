package db

import (
	"gorm.io/gorm"
	"time"
)

type Conf struct {
	Dsn         string
	MaxIdle     int
	MaxOpen     int
	MaxLifeTime int
}

func NewDb(conf *Conf) *gorm.DB {

	instance := New()

	if err := instance.Connect(conf.Dsn, &gorm.Config{}); err != nil {
		panic(err)
	}

	if err := instance.EnableConnectionPool(conf.MaxIdle, conf.MaxOpen, time.Duration(conf.MaxLifeTime)); err != nil {
		panic(err)
	}

	return instance.DB
}
