package main

import (
	"encoding/json"
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
	order,err := order.CreateOrder(100000,1)
	if err != nil{
		w.Write([]byte(err.Error()))
	}
	b,_ := json.Marshal(order)
	w.Write(b)
}
