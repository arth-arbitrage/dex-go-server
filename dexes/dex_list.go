package dexes

import (
	"github.com/arth-arbitrage/dex-go-server/go"
)

type DexIf interface {
	Name(ctx *DefaultApiService) string
	GetSwapPools(ctx *DefaultApiService) (restapi.ImplResponse, error)
	Init(ctx *DefaultApiService) error
}

var dexList = []DexIf {
	&DexUniswapV2{name: "UniswapV2"},
	&DexCurve{name: "Curve"},	
}
