package order

import (
	"errors"
	"miaosha/Dao"
	"miaosha/common/db"
)

type Order struct{
}

//创建订单
func CreateOrder(oid int64,nums int32)(order *db.Order,err error){
	var(
		goods *db.Goods
	)

	//商品是否存在
	goods,err = Dao.GetGoods(oid)
	if err != nil{
		return
	}

	//检查库存
	if int32(goods.Stocknum) < nums{
		err = errors.New("商品库存不足")
		return
	}

	//开启事务
	session := db.Engine().NewSession()



}
