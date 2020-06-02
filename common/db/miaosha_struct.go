package db

import (
	"time"
)

// Goods -
type Goods struct {
	Oid         int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Name        string    `xorm:"VARCHAR(255)"`
	Price       int64     `xorm:"default 0 BIGINT(20)"`
	Stocknum    int       `xorm:"default 0 INT(11)"`
	Description string    `xorm:"VARCHAR(255)"`
	Status      int       `xorm:"default 1 TINYINT(4)"`
	Createtime  time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// Order -
type Order struct {
	Orderid    int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Orderno    string    `xorm:"VARCHAR(64)"`
	Soid       int64     `xorm:"default 0 BIGINT(20)"`
	Oid        int64     `xorm:"default 0 BIGINT(20)"`
	Price      int64     `xorm:"default 0 BIGINT(20)"`
	Num        int       `xorm:"default 0 INT(11)"`
	Status     int       `xorm:"default 0 TINYINT(4)"`
	Payuid     int64     `xorm:"default 0 BIGINT(20)"`
	Paystatus  int       `xorm:"default 0 TINYINT(4)"`
	Paytime    time.Time `xorm:"TIMESTAMP"`
	Mobile     string    `xorm:"VARCHAR(24)"`
	Address    string    `xorm:"VARCHAR(255)"`
	Createtime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Updatetime time.Time `xorm:"TIMESTAMP"`
}

// User -
type User struct {
	Uid        int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Nick       string    `xorm:"VARCHAR(255)"`
	Password   string    `xorm:"VARCHAR(32)"`
	Logo       string    `xorm:"VARCHAR(255)"`
	Createtime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Logintime  time.Time `xorm:"TIMESTAMP"`
}
