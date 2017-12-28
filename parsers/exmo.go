package parsers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type CourceExmo struct {
	BuyPrice  float64 `json:"buy_price"`
	SellPrice float64 `json:"sell_price"`
	Course    float64 `json:"-"`
	// LastTrade float32 `json:"last_trade"`
	// High      float32 `json:"high"`
	// Low       float32 `json:"low"`
	// Avg float32 `json:"avg,string"`
	// Vol       float32 `json:"vol"`
	// VolCurr   float32 `json:"vol_curr"`
	// Updated   uint32  `json:"updated"`
}

//easyjson:json
type Exmo struct {
	BtcUsd   CourceExmo `json:"BTC_USD"`
	BtcEur   CourceExmo `json:"BTC_EUR"`
	BtcRub   CourceExmo `json:"BTC_RUB"`
	BtcUah   CourceExmo `json:"BTC_UAH"`
	BtcPln   CourceExmo `json:"BTC_PLN"`
	BchBtc   CourceExmo `json:"BCH_BTC"`
	BchUsd   CourceExmo `json:"BCH_USD"`
	BchRub   CourceExmo `json:"BCH_RUB"`
	BchEth   CourceExmo `json:"BCH_ETH"`
	DashBtc  CourceExmo `json:"DASH_BTC"`
	DashUsd  CourceExmo `json:"DASH_USD"`
	DashRub  CourceExmo `json:"DASH_RUB"`
	EthBtc   CourceExmo `json:"ETH_BTC"`
	EthLtc   CourceExmo `json:"ETH_LTC"`
	EthUsd   CourceExmo `json:"ETH_USD"`
	EthEur   CourceExmo `json:"ETH_EUR"`
	EthRub   CourceExmo `json:"ETH_RUB"`
	EthUah   CourceExmo `json:"ETH_UAH"`
	EthPln   CourceExmo `json:"ETH_PLN"`
	EtcBtc   CourceExmo `json:"ETC_BTC"`
	EtcUsd   CourceExmo `json:"ETC_USD"`
	EtcRub   CourceExmo `json:"ETC_RUB"`
	LtcBtc   CourceExmo `json:"LTC_BTC"`
	LtcUsd   CourceExmo `json:"LTC_USD"`
	LtcEur   CourceExmo `json:"LTC_EUR"`
	LtcRub   CourceExmo `json:"LTC_RUB"`
	ZecBtc   CourceExmo `json:"ZEC_BTC"`
	ZecUsd   CourceExmo `json:"ZEC_USD"`
	ZecEur   CourceExmo `json:"ZEC_EUR"`
	ZecRub   CourceExmo `json:"ZEC_RUB"`
	XrpBtc   CourceExmo `json:"XRP_BTC"`
	XrpUsd   CourceExmo `json:"XRP_USD"`
	XrpRub   CourceExmo `json:"XRP_RUB"`
	XmrBtc   CourceExmo `json:"XMR_BTC"`
	XmrUsd   CourceExmo `json:"XMR_USD"`
	XmrEur   CourceExmo `json:"XMR_EUR"`
	BtcUsdt  CourceExmo `json:"BTC_USDT"`
	EthUsdt  CourceExmo `json:"ETH_USDT"`
	UsdtUsd  CourceExmo `json:"USDT_USD"`
	UsdtRub  CourceExmo `json:"USDT_RUB"`
	UsdRub   CourceExmo `json:"USD_RUB"`
	DogeBtc  CourceExmo `json:"DOGE_BTC"`
	WavesBtc CourceExmo `json:"WAVES_BTC"`
	WavesRub CourceExmo `json:"WAVES_RUB"`
	KickBtc  CourceExmo `json:"KICK_BTC"`
	KickEth  CourceExmo `json:"KICK_ETH"`
}

var courceExmo CourceExmo

func LoadExmoData() (CourceExmo, error) {
	_, raw, err := fasthttp.Get(nil, "https://api.exmo.com/v1/ticker/")

	if nil != err {
		return courceExmo, err
	}

	return ParseExmoData(raw)
}

func ParseExmoData(data []byte) (CourceExmo, error) {
	return courceExmo, json.Unmarshal(data, &courceExmo)
}
