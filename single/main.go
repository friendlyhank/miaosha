package main

import (
	_ "miaosha/common"
	"miaosha/single/order"
	"net/http"
)

func main(){
	http.HandleFunc("/buy/goods",HandleBuyGoodsfunc)
	http.ListenAndServe(":8030",nil)
}

//HandleBuyGoodsfunc- 购买商品
func HandleBuyGoodsfunc(w http.ResponseWriter,r *http.Request){
	order.CreateOrder(100000,1)
	w.Write([]byte("恭喜你,秒杀成功!"))
}
