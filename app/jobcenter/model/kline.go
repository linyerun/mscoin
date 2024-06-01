package model

import "mscoin/common/tool"

type Kline struct {
	Period       string  `bson:"period,omitempty"`
	OpenPrice    float64 `bson:"openPrice,omitempty"`
	HighestPrice float64 `bson:"highestPrice,omitempty"`
	LowestPrice  float64 `bson:"lowestPrice,omitempty"`
	ClosePrice   float64 `bson:"closePrice,omitempty"`
	Time         int64   `bson:"time,omitempty"`
	Count        float64 `bson:"count,omitempty"`    //成交笔数
	Volume       float64 `bson:"volume,omitempty"`   //成交量
	Turnover     float64 `bson:"turnover,omitempty"` //成交额
	TimeStr      string  `bson:"timeStr,omitempty"`
}

func NewKline(data []string, period string) *Kline {
	toInt64 := tool.ToInt64(data[0])
	return &Kline{
		Time:         toInt64,
		Period:       period,
		OpenPrice:    tool.ToFloat64(data[1]),
		HighestPrice: tool.ToFloat64(data[2]),
		LowestPrice:  tool.ToFloat64(data[3]),
		ClosePrice:   tool.ToFloat64(data[4]),
		Count:        tool.ToFloat64(data[5]),
		Volume:       tool.ToFloat64(data[6]),
		Turnover:     tool.ToFloat64(data[7]),
		TimeStr:      tool.ToTimeString(toInt64),
	}
}

// GetKlineTableName BTC/USDT  ETH/USDT  分表
func GetKlineTableName(symbol, period string) string {
	return "exchange_kline_" + symbol + "_" + period
}
