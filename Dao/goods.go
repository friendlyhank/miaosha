package Dao

import (
	"github.com/xormplus/xorm"
	"miaosha/common/db"
)

//GetGoods-获取商品
func GetGoods(oid int64)(goods *db.Goods,err error){
	goods = &db.Goods{}
	var has bool
	has,err = db.Engine().ID(oid).Get(goods)
	if !has || err != nil{
		goods = nil
	}
	return
}

//修改商品
func UpdateGoods(session *xorm.Session,goods *db.Goods,columns ...string)(int64,error){
	return session.ID(goods.Oid).Cols(columns...).Update(goods)
}
