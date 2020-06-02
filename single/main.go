package main

import "net/http"

func main(){
	http.HandleFunc("/buy/goods",HandleBuyGoodsfunc)
	http.ListenAndServe(":8030",nil)
}

//HandleBuyGoodsfunc- 购买商品
func HandleBuyGoodsfunc(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("你好"))
}
