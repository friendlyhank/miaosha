package order

import (
	_ "miaosha/common"
	"testing"
)


func TestCreateOrder(t *testing.T){
	CreateOrder(100000,10)
}
