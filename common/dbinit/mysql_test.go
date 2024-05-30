package dbinit

import (
	"fmt"
	"testing"
)

func TestMySQLConnect(t *testing.T) {
	db, err := CreateGormMysqlClient("root:root@tcp(192.168.200.143:3309)/mscoin?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", 10, 100)
	if err != nil {
		t.Fatal(err)
	}

	type Coin struct {
		Id int64 `gorm:"column:id"`
	}

	coin := new(Coin)
	err = db.Table("coin").Select("id").Limit(1).First(coin).Error
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(coin)
}
