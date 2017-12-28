package parsers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type CourcePoloniex struct {
	// ID   uint8   `json:"id"`
	// Last float32 `json:"last,string"`
	LowestAsk  float32 `json:"lowestAsk,string"`
	HighestBid float32 `json:"highestBid,string"`
	// PercentChange float32 `json:"percentChange,string"`
	// BaseVolume    float32 `json:"baseVolume,string"`
	// QuoteVolume   float32 `json:"quoteVolume,string"`
	// IsFrozen      uint8   `json:"isFrozen,string"`
	// High24hr      float32 `json:"high24hr,string"`
	// Low24hr       float32 `json:"low24hr,string"`
}

//easyjson:json
type Poloniex struct {
	BtcBcn   CourcePoloniex `json:"BTC_BCN"`
	BtcBela  CourcePoloniex `json:"BTC_BELA"`
	BtcBlk   CourcePoloniex `json:"BTC_BLK"`
	BtcBtcd  CourcePoloniex `json:"BTC_BTCD"`
	BtcBtm   CourcePoloniex `json:"BTC_BTM"`
	BtcBts   CourcePoloniex `json:"BTC_BTS"`
	BtcBurst CourcePoloniex `json:"BTC_BURST"`
	BtcClam  CourcePoloniex `json:"BTC_CLAM"`
	BtcDash  CourcePoloniex `json:"BTC_DASH"`
	BtcDgb   CourcePoloniex `json:"BTC_DGB"`
	BtcDoge  CourcePoloniex `json:"BTC_DOGE"`
	BtcEmc2  CourcePoloniex `json:"BTC_EMC2"`
	BtcFldc  CourcePoloniex `json:"BTC_FLDC"`
	BtcFlo   CourcePoloniex `json:"BTC_FLO"`
	BtcGame  CourcePoloniex `json:"BTC_GAME"`
	BtcGrc   CourcePoloniex `json:"BTC_GRC"`
	BtcHuc   CourcePoloniex `json:"BTC_HUC"`
	BtcLtc   CourcePoloniex `json:"BTC_LTC"`
	BtcMaid  CourcePoloniex `json:"BTC_MAID"`
	BtcOmni  CourcePoloniex `json:"BTC_OMNI"`
	BtcNav   CourcePoloniex `json:"BTC_NAV"`
	BtcNeos  CourcePoloniex `json:"BTC_NEOS"`
	BtcNmc   CourcePoloniex `json:"BTC_NMC"`
	BtcNxt   CourcePoloniex `json:"BTC_NXT"`
	BtcPink  CourcePoloniex `json:"BTC_PINK"`
	BtcPot   CourcePoloniex `json:"BTC_POT"`
	BtcPpc   CourcePoloniex `json:"BTC_PPC"`
	BtcRic   CourcePoloniex `json:"BTC_RIC"`
	BtcStr   CourcePoloniex `json:"BTC_STR"`
	BtcSys   CourcePoloniex `json:"BTC_SYS"`
	BtcVia   CourcePoloniex `json:"BTC_VIA"`
	BtcXvc   CourcePoloniex `json:"BTC_XVC"`
	BtcVrc   CourcePoloniex `json:"BTC_VRC"`
	BtcVtc   CourcePoloniex `json:"BTC_VTC"`
	BtcXbc   CourcePoloniex `json:"BTC_XBC"`
	BtcXcp   CourcePoloniex `json:"BTC_XCP"`
	BtcXem   CourcePoloniex `json:"BTC_XEM"`
	BtcXmr   CourcePoloniex `json:"BTC_XMR"`
	BtcXpm   CourcePoloniex `json:"BTC_XPM"`
	BtcXrp   CourcePoloniex `json:"BTC_XRP"`
	UsdtBtc  CourcePoloniex `json:"USDT_BTC"`
	UsdtDash CourcePoloniex `json:"USDT_DASH"`
	UsdtLtc  CourcePoloniex `json:"USDT_LTC"`
	UsdtNxt  CourcePoloniex `json:"USDT_NXT"`
	UsdtStr  CourcePoloniex `json:"USDT_STR"`
	UsdtXmr  CourcePoloniex `json:"USDT_XMR"`
	UsdtXrp  CourcePoloniex `json:"USDT_XRP"`
	XmrBcn   CourcePoloniex `json:"XMR_BCN"`
	XmrBlk   CourcePoloniex `json:"XMR_BLK"`
	XmrBtcd  CourcePoloniex `json:"XMR_BTCD"`
	XmrDash  CourcePoloniex `json:"XMR_DASH"`
	XmrLtc   CourcePoloniex `json:"XMR_LTC"`
	XmrMaid  CourcePoloniex `json:"XMR_MAID"`
	XmrNxt   CourcePoloniex `json:"XMR_NXT"`
	BtcEth   CourcePoloniex `json:"BTC_ETH"`
	UsdtEth  CourcePoloniex `json:"USDT_ETH"`
	BtcSc    CourcePoloniex `json:"BTC_SC"`
	BtcBcy   CourcePoloniex `json:"BTC_BCY"`
	BtcExp   CourcePoloniex `json:"BTC_EXP"`
	BtcFct   CourcePoloniex `json:"BTC_FCT"`
	BtcRads  CourcePoloniex `json:"BTC_RADS"`
	BtcAmp   CourcePoloniex `json:"BTC_AMP"`
	BtcDcr   CourcePoloniex `json:"BTC_DCR"`
	BtcLsk   CourcePoloniex `json:"BTC_LSK"`
	EthLsk   CourcePoloniex `json:"ETH_LSK"`
	BtcLbc   CourcePoloniex `json:"BTC_LBC"`
	BtcSteem CourcePoloniex `json:"BTC_STEEM"`
	EthSteem CourcePoloniex `json:"ETH_STEEM"`
	BtcSbd   CourcePoloniex `json:"BTC_SBD"`
	BtcEtc   CourcePoloniex `json:"BTC_ETC"`
	EthEtc   CourcePoloniex `json:"ETH_ETC"`
	UsdtEtc  CourcePoloniex `json:"USDT_ETC"`
	BtcRep   CourcePoloniex `json:"BTC_REP"`
	UsdtRep  CourcePoloniex `json:"USDT_REP"`
	EthRep   CourcePoloniex `json:"ETH_REP"`
	BtcArdr  CourcePoloniex `json:"BTC_ARDR"`
	BtcZec   CourcePoloniex `json:"BTC_ZEC"`
	EthZec   CourcePoloniex `json:"ETH_ZEC"`
	UsdtZec  CourcePoloniex `json:"USDT_ZEC"`
	XmrZec   CourcePoloniex `json:"XMR_ZEC"`
	BtcStrat CourcePoloniex `json:"BTC_STRAT"`
	BtcNxc   CourcePoloniex `json:"BTC_NXC"`
	BtcPasc  CourcePoloniex `json:"BTC_PASC"`
	BtcGnt   CourcePoloniex `json:"BTC_GNT"`
	EthGnt   CourcePoloniex `json:"ETH_GNT"`
	BtcGno   CourcePoloniex `json:"BTC_GNO"`
	EthGno   CourcePoloniex `json:"ETH_GNO"`
	BtcBch   CourcePoloniex `json:"BTC_BCH"`
	EthBch   CourcePoloniex `json:"ETH_BCH"`
	UsdtBch  CourcePoloniex `json:"USDT_BCH"`
	BtcZrx   CourcePoloniex `json:"BTC_ZRX"`
	EthZrx   CourcePoloniex `json:"ETH_ZRX"`
	BtcCvc   CourcePoloniex `json:"BTC_CVC"`
	EthCvc   CourcePoloniex `json:"ETH_CVC"`
	BtcOmg   CourcePoloniex `json:"BTC_OMG"`
	EthOmg   CourcePoloniex `json:"ETH_OMG"`
	BtcGas   CourcePoloniex `json:"BTC_GAS"`
	EthGas   CourcePoloniex `json:"ETH_GAS"`
	BtcStorj CourcePoloniex `json:"BTC_STORJ"`
}

var courcePoloniex CourcePoloniex

func LoadPoloniexData() (CourcePoloniex, error) {
	_, raw, err := fasthttp.Get(nil, "https://poloniex.com/public?command=returnTicker")

	if nil != err {
		return courcePoloniex, err
	}

	return ParsePoloniexData(raw)
}

func ParsePoloniexData(data []byte) (CourcePoloniex, error) {
	return courcePoloniex, json.Unmarshal(data, &courcePoloniex)
}
