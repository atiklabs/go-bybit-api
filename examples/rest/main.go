package main

import (
	"github.com/atiklabs/go-bybit-api/rest"
	"log"
)

func main() {
	//baseURL := "https://api.bybit.com/"	// 主网络
	baseURL := "https://api-testnet.bybit.com/" // 测试网络
	b := rest.New(nil,
		baseURL, "YIxOY2RhFkylPudq96", "Bg9G2oFOb3aaIMguD3FOvOJJVBycaoXqXNcI", true)

	// 获取持仓
	_, _, positions, err := b.GetPositions()
	if err != nil {
		log.Printf("%v", err)
		return
	}

	log.Printf("positions: %#v", positions)

	// 创建委托
	symbol := "BTCUSD"
	side := "Buy"
	orderType := "Limit"
	qty := 30
	price := 7000.0
	timeInForce := "GoodTillCancel"
	_, _, order, err := b.CreateOrder(side, orderType, price, qty, timeInForce, 0, 0, false, false, "", symbol)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Create order: %#v", order)

	// 获取委托单
}
