package db

import (
	"git.biezao.com/ant/xmiss/foundation/vars"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	mysql "github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

var (
	// miaosha 的数据库
	miaoshaEngine *xorm.Engine
)

// Init - 自动初始化数据库的链接
func Init() {
	logs.Debug("|foundation|init|db|Init")

	var (
		err error
	)
	// 把这个依赖关系放到这里, 这样这里其实是初始化了驱动为Mysql
	if mysql.ErrInvalidConn != nil {
	}

	// miaosha
	dbSource := beego.AppConfig.String("dbsource")
	if miaoshaEngine, err = xorm.NewEngine("mysql", dbSource); err != nil {
		logs.Error("Engine Init Err:%v", err)
		panic(err)
	}
	// 获取当前的运行模式，如果再DEV模式下，则打印所有的Sql
	if !vars.IsProd() {
		miaoshaEngine.ShowSQL(true)
		miaoshaEngine.ShowExecTime(true)
		miaoshaEngine.SetLogger(xorm.NewSimpleLogger(beego.BeeLogger))
		miaoshaEngine.Logger().SetLevel(core.LOG_INFO)
	}
}

// Engine - 主数据库
func Engine() *xorm.Engine {
	return miaoshaEngine
}
