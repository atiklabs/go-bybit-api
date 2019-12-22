package main

import (
	"github.com/frankrap/bybit-api/ws"
	"log"
)

func main() {
	cfg := &ws.Configuration{
		Addr:          ws.HostTestnet, // 测试网络
		ApiKey:        "wKuYtkeNdC2PaMKjoy",
		SecretKey:     "5ekcDn3KnKoCRbfvrPImYzVdx7Ri2hhVxkmw",
		AutoReconnect: true, // 断线自动重连
		DebugMode:     true,
	}
	b := ws.New(cfg)

	// 订阅新版25档orderBook
	b.Subscribe(ws.WSOrderBook25L1 + ".BTCUSD")
	// 实时交易
	//b.Subscribe("trade.BTCUSD")
	b.Subscribe(ws.WSTrade) // BTCUSD/ETHUSD/EOSUSD/XRPUSD
	// K线
	b.Subscribe(ws.WSKLine + ".BTCUSD.1m")
	// 每日保险基金更新
	b.Subscribe(ws.WSInsurance)
	// 产品最新行情
	b.Subscribe(ws.WSInstrument + ".BTCUSD")

	// 仓位变化
	b.Subscribe(ws.WSPosition)
	// 委托单成交信息
	b.Subscribe(ws.WSExecution)
	// 委托单的更新
	b.Subscribe(ws.WSOrder)

	b.On(ws.WSOrderBook25L1, handleOrderBook)
	b.On(ws.WSTrade, handleTrade)
	b.On(ws.WSKLine, handleKLine)
	b.On(ws.WSInsurance, handleInsurance)
	b.On(ws.WSInstrument, handleInstrument)

	b.On(ws.WSPosition, handlePosition)
	b.On(ws.WSExecution, handleExecution)
	b.On(ws.WSOrder, handleOrder)

	b.Start()

	forever := make(chan struct{})
	<-forever
}

func handleOrderBook(symbol string, data ws.OrderBook) {
	log.Printf("handleOrderBook %v/%v", symbol, data)
}

func handleTrade(symbol string, data []*ws.Trade) {
	log.Printf("handleTrade %v/%v", symbol, data)
}

func handleKLine(symbol string, data ws.KLine) {
	log.Printf("handleKLine %v/%v", symbol, data)
}

func handleInsurance(currency string, data []*ws.Insurance) {
	log.Printf("handleInsurance %v/%v", currency, data)
}

func handleInstrument(symbol string, data []*ws.Instrument) {
	log.Printf("handleInstrument %v/%v", symbol, data)
}

func handlePosition(data []*ws.Position) {
	log.Printf("handlePosition %v", data)
}

func handleExecution(data []*ws.Execution) {
	log.Printf("handleExecution %v", data)
}

func handleOrder(data []*ws.Order) {
	log.Printf("handleOrder %v", data)
}
