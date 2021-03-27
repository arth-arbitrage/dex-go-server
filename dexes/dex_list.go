package dexes

import (
	//"context"
	"github.com/arth-arbitrage/dex-go-server/go"
)

type DexIf interface {
	Name() string
	GetSwapPools(ctx *DefaultApiService) (restapi.ImplResponse, error)
}

var dexList = []DexIf {
	&DexUniswapV2{name: "UniswapV2"},
	&DexCurve{name: "Curve"},	
}
/*
func AddDex(dex interface) {
	dexList = append(dexList, dex)
}
*/