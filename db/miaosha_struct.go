package db

import (
	"time"
)

// Goods -
type Goods struct {
	Oid int64 `xorm:"not null pk autoincr BIGINT(20)"`
}

// Order -
type Order struct {
	Orderid    int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Orderno    string    `xorm:"VARCHAR(64)"`
	Soid       int64     `xorm:"BIGINT(20)"`
	Oid        int64     `xorm:"BIGINT(20)"`
	Price      int64     `xorm:"BIGINT(20)"`
	Num        int       `xorm:"INT(11)"`
	Status     int       `xorm:"TINYINT(4)"`
	Payuid     int64     `xorm:"BIGINT(20)"`
	Paystatus  int       `xorm:"TINYINT(4)"`
	Paytime    time.Time `xorm:"TIMESTAMP"`
	Mobile     string    `xorm:"VARCHAR(24)"`
	Address    string    `xorm:"VARCHAR(255)"`
	Createtime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Updatetime time.Time `xorm:"TIMESTAMP"`
}
