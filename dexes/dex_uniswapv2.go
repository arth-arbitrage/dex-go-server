package dexes

import (
	"net/http"
	"fmt"
	"math"
	//"errors"
	//"context"
	"math/big"	
	//"encoding/json"
	"github.com/arth-arbitrage/dex-go-server/go-contracts/uniswapv2"
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"	
	"github.com/arth-arbitrage/dex-go-server/go"
)

type DexUniswapV2 struct {
	name string
}

type token struct {
	name string
	address  common.Address
	decimals int
}



var tokenList = []token {
	token {name: "WETH", address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), decimals:18 },
    token {name: "DAI", address:common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"), decimals:18},
    token {name: "USDC", address: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"), decimals:6},
    token {name: "USDT", address: common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"), decimals:6},
    token {name: "wBTC", address: common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"), decimals:8},
    token {name: "renBTC", address: common.HexToAddress("0xeb4c2781e4eba804ce9a9803c67d0893436bb27d"), decimals:8},
    token {name: "BUSD", address: common.HexToAddress("0x4fabb145d64652a948d72533023f6e7a623c7c53"), decimals:18},
    token {name: "TUSD", address: common.HexToAddress("0x0000000000085d4780b73119b644ae5ecd22b376"), decimals:18},
    token {name: "EURS", address: common.HexToAddress("0xdB25f211AB05b1c97D595516F45794528a807ad8"), decimals:2},
    token {name: "VID", address: common.HexToAddress("0x2c9023bbc572ff8dc1228c7858a280046ea8c9e5"), decimals:18},
}

const factoryAddress string = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
const routerAddress string = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"

func NewDexUniswapV2() *DexUniswapV2 {
	return &DexUniswapV2{ name: "uniswapv2"}
}

func (dex DexUniswapV2) Name() string {
	return "UniswapV2"
}

func (dex DexUniswapV2)GetSwapPools(ctx *DefaultApiService) (restapi.ImplResponse, error) {
//func (dex DexUniswapV2) GetSwapPools(w http.ResponseWriter, r *http.Request) {
	client := ctx.client

	factory, err := uniswapv2.NewIUniswapV2Factory(common.Address(common.HexToAddress(factoryAddress)), client)
	if (err != nil) { fmt.Printf("Failed to get factory at %v",factoryAddress)}

	//router, err := uniswapv2.NewIUniswapV2Router02(common.Address(common.HexToAddress(routerAddress)), client)
	swapPools := make([]restapi.SwapPool, 0)
	for i, token0 :=  range tokenList {
		for _, token1 :=  range tokenList[i+1:] {
			if (token0 != token1) {
				fmt.Printf("Pair for %v %v\n", token0.name, token1.name)
				opts  := bind.CallOpts{Pending: false}
				pairAddr, err := factory.GetPair(&opts, token0.address, token1.address)
				if (err != nil || pairAddr == common.Address(common.HexToAddress("x0"))) { 
					fmt.Printf("Failed to get pair for %v %v", token0.name, token1.name)
					continue
				}
				
				pair, err := uniswapv2.NewIUniswapV2Pair(pairAddr, client)
				if (err != nil) { 
					fmt.Printf("Failed to create pair for %v %v", token0.name, token1.name)
					continue
				} 
			
				reserves, err := pair.GetReserves(&opts)
				if err != nil {
					fmt.Printf("Failed to get pair data for %v %v", token0.name, token1.name)
				} else {
					token0Addr,  _ := pair.Token0(&opts)
					from_token, to_token := token0, token1 
					if(token0.address != token0Addr) {
						from_token, to_token = token1, token0 
					}
					
					t0_div := new(big.Float).SetFloat64(math.Pow(10, float64(from_token.decimals)))
					t1_div := new(big.Float).SetFloat64(math.Pow(10, float64(to_token.decimals)))

					f0 := new(big.Float).SetInt(reserves.Reserve0)
					f0 = new(big.Float).Quo(f0, t0_div)
					f0_f32, _ := f0.Float32()

					f1 := new(big.Float).SetInt(reserves.Reserve1)
					f1 = new(big.Float).Quo(f1, t1_div)
					f1_f32, _ := f1.Float32()

					p0_f32 := f1_f32/f0_f32
					p1_f32 := f0_f32/f1_f32

					swapPool := restapi.SwapPool{
						Id: 0, 
						Name:"token0",
						Token0: from_token.name,
						Token1: to_token.name,
						Reserve0: f0_f32,
						Reserve1: f1_f32,
						Price0:p0_f32,
						Price1:p1_f32,
						Fees:0.003,
					}

					swapPools = append(swapPools, swapPool)
				}
			}
		}
	}
	//jsonData, err := json.Marshal(swapPools)
	//if err != nil {
	//	return restapi.Response(http.StatusInternalServerError, nil), err
	//}	
	return restapi.Response(http.StatusOK, swapPools/*(string(jsonData)*/), nil	
}