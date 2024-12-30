package types

type Rates struct {
	Rates coinRates `json:"rates"`
}

type rateInfo struct {
	Name  string  `json:"name"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

type coinRates struct {
	Btc  rateInfo `json:"btc"`
	Eth  rateInfo `json:"eth"`
	Ltc  rateInfo `json:"ltc"`
	Bch  rateInfo `json:"bch"`
	Bnb  rateInfo `json:"bnb"`
	Eos  rateInfo `json:"eos"`
	Xrp  rateInfo `json:"xrp"`
	Xlm  rateInfo `json:"xlm"`
	Link rateInfo `json:"link"`
	Dot  rateInfo `json:"dot"`
	Yfi  rateInfo `json:"yfi"`
	Usd  rateInfo `json:"usd"`
	Aed  rateInfo `json:"aed"`
	Ars  rateInfo `json:"ars"`
	Aud  rateInfo `json:"aud"`
	Bdt  rateInfo `json:"bdt"`
	Bhd  rateInfo `json:"bhd"`
	Bmd  rateInfo `json:"bmd"`
	Brl  rateInfo `json:"brl"`
	Cad  rateInfo `json:"cad"`
	Chf  rateInfo `json:"chf"`
	Clp  rateInfo `json:"clp"`
	Cny  rateInfo `json:"cny"`
	Czk  rateInfo `json:"czk"`
	Dkk  rateInfo `json:"dkk"`
	Eur  rateInfo `json:"eur"`
	Gbp  rateInfo `json:"gbp"`
	Hkd  rateInfo `json:"hkd"`
	Huf  rateInfo `json:"huf"`
	Idr  rateInfo `json:"idr"`
	Ils  rateInfo `json:"ils"`
	Inr  rateInfo `json:"inr"`
	Jpy  rateInfo `json:"jpy"`
	Krw  rateInfo `json:"krw"`
	Kwd  rateInfo `json:"kwd"`
	Lkr  rateInfo `json:"lkr"`
	Mmk  rateInfo `json:"mmk"`
	Mxn  rateInfo `json:"mxn"`
	Myr  rateInfo `json:"myr"`
	Ngn  rateInfo `json:"ngn"`
	Nok  rateInfo `json:"nok"`
	Nzd  rateInfo `json:"nzd"`
	Php  rateInfo `json:"php"`
	Pkr  rateInfo `json:"pkr"`
	Pln  rateInfo `json:"pln"`
	Rub  rateInfo `json:"rub"`
	Sar  rateInfo `json:"sar"`
	Sek  rateInfo `json:"sek"`
	Sgd  rateInfo `json:"sgd"`
	Thb  rateInfo `json:"thb"`
	Try  rateInfo `json:"try"`
	Twd  rateInfo `json:"twd"`
	Uah  rateInfo `json:"uah"`
	Vef  rateInfo `json:"vef"`
	Vnd  rateInfo `json:"vnd"`
	Zar  rateInfo `json:"zar"`
	Xdr  rateInfo `json:"xdr"`
	Xag  rateInfo `json:"xag"`
	Xau  rateInfo `json:"xau"`
	Bits rateInfo `json:"bits"`
	Sats rateInfo `json:"sats"`
}
