package dexes

import (
	"github.com/arth-arbitrage/dex-go-server/go"
)

type LendIf interface {
	Name(ctx *DefaultApiService) string
	GetLenderPools(ctx *DefaultApiService) (restapi.ImplResponse, error)
	Init(ctx *DefaultApiService) error
}

var lendList = []LendIf {
	&LendAave{name:"aave"},
	&LendArth{name:"arth"},	
}