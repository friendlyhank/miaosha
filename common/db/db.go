package db

import (
	mysql "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var (
	// miaosha 的数据库
	miaoshaEngine *xorm.Engine
)

// Init - 自动初始化数据库的链接
func Init() {
	var (
		err error
	)
	// 把这个依赖关系放到这里, 这样这里其实是初始化了驱动为Mysql
	if mysql.ErrInvalidConn != nil {
	}

	// miaosha
	dbSource := "root:123456@tcp(127.0.0.1:3306)/miaosha?charset=utf8mb4&loc=Local&interpolateParams=true"
	if miaoshaEngine, err = xorm.NewEngine("mysql", dbSource); err != nil {
		panic(err)
	}
	// 获取当前的运行模式，如果再DEV模式下，则打印所有的Sql
		miaoshaEngine.ShowSQL(true)
}

// Engine - 主数据库
func Engine() *xorm.Engine {
	return miaoshaEngine
}
