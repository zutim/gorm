package db

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	db := NewDb(&Conf{
		Dsn:         "user:password@tcp(127.0.0.1:3306)/onebet_sport?charset=utf8mb4&parseTime=True&loc=Local",
		MaxOpen:     4,
		MaxIdle:     10,
		MaxLifeTime: 100,
	})

	type Tickets struct {
		Id string
	}

	var ti Tickets

	if err := db.Table("tbl_tickets").First(&ti).Error; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ti.Id)
}
