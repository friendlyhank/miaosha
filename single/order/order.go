package order

import (
	"errors"
	"miaosha/Dao"
	"miaosha/common/db"
	"time"
)

type Order struct{
}

//创建订单
func CreateOrder(oid int64,nums int)(order *db.Order,err error){

	//商品是否存在
	var g *db.Goods
	g,err = Dao.GetGoods(oid)
	if err != nil{
		return
	}

	if g == nil{
		err = errors.New("商品不存在")
		return
	}

	//检查库存
	if g.Stocknum < nums{
		err = errors.New("商品库存不足")
		return
	}

	session := db.Engine().NewSession()
	defer session.Close()

	//开启事务
	session.Begin()

	//排它锁锁定订单
	var goods = &db.Goods{}
	var has bool
	has,err = session.Table("goods").Where("oid =?",g.Oid).ForUpdate().Get(goods)
	if !has || err != nil{
		session.Rollback()
		err = errors.New("商品不存在")
		return
	}

	//检查库存
	if goods.Stocknum < nums{
		session.Rollback()
		err = errors.New("商品库存不足")
		return
	}

	//扣减库存
	goods.Stocknum -= int(nums)
	_,err = Dao.UpdateGoods(session,goods,"stocknum")
	if err != nil{
		session.Rollback()
		return
	}

	//生成订单
	order = &db.Order{
		Orderid:    0,
		Orderno:    "",
		Soid:       1,
		Oid:        goods.Oid,
		Price:      goods.Price,
		Num:        nums,
		Status:     1,
		Payuid:     1,
		Paystatus:  1,
		Paytime:    time.Now(),
		Mobile:     "16920145897",
		Address:    "广东广州",
		Createtime: time.Now(),
		Updatetime: time.Now(),
	}
	_,err = session.Insert(order)
	if err != nil{
		session.Rollback()
		return
	}

	session.Commit()
	return
}
