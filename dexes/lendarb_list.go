package dexes

import (
	"github.com/arth-arbitrage/dex-go-server/go"
)

type MultiswapIf interface {
	Name(ctx *DefaultApiService) string
	Multiswap(ctx *DefaultApiService, multiSwap restapi.MultiSwap) (restapi.ImplResponse, error)
	Init(ctx *DefaultApiService) error
}

var lendarbList = []MultiswapIf {
	&ArthLendArb{name:"arth"},
	&AaveLendArb{name:"aave"},	
}