package main

import (
	"fmt"
	"strconv"
)

/* 接口与多态 */
// 定义一个商品接口
type Good interface {
	settleAccount() int // 金额
	orderInfo() string  // 商品信息
}

// 定义一个电话结构体
type Phone struct {
	name     string // 商品名称
	quantity int    // 数量
	price    int    // 单价
}

// 实现接口方法, 返回商品价格
func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}

// 实现接口方法，返回商品信息及价格信息
func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计: " + strconv.Itoa(phone.settleAccount()) + "元"
}

// 定义一个赠品结构体
type FreeGift struct {
	name     string // 赠品名称
	quantity int    // 赠品数量
	price    int    // 赠品单价
}

// 实现接口方法, 返回商品价格
func (freeGift FreeGift) settleAccount() int {
	return 0
}

// 实现接口方法，返回商品信息及价格信息
func (freeGift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(freeGift.quantity) + "个" +
		freeGift.name + "计: " + strconv.Itoa(freeGift.settleAccount()) + "元"
}

// 定义一个方法来计算购物车中的金额
func calculateAllPrice(goods []Good) int {
	var allPrice int // 定义总金额
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}

// main()方法
func main() {
	iphone := Phone{
		name:     "iPhone12 Pro",
		quantity: 1,
		price:    7999,
	}
	airPods := FreeGift{
		name:     "airPods",
		quantity: 1,
		price:    2599,
	}
	macBookPro := Phone{
		name:     "MacBook Pro",
		quantity: 1,
		price:    12999,
	}

	// 创建一个购物车，类型为Good的切片，存放相关商品
	goods := []Good{iphone, airPods, macBookPro}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单共需支付 %d 元\n", allPrice)
}
