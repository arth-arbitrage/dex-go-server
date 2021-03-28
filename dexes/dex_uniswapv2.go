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
	factory *uniswapv2.IUniswapV2Factory
	routerv2 *uniswapv2.IUniswapV2Router02
}

const uniswapv2FactoryAddress string = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
const uniswapv2RouterAddress string = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"

func (dex DexUniswapV2) Name(ctx *DefaultApiService) string {
	return dex.name
}

func (dex *DexUniswapV2) Init(ctx *DefaultApiService) error {
	var err error
	client := ctx.client
	dex.factory, err = uniswapv2.NewIUniswapV2Factory(common.Address(common.HexToAddress(uniswapv2FactoryAddress)), client)
	if (err != nil) { 
		fmt.Printf("Failed to get factory at %v",uniswapv2FactoryAddress)
		return err
	}
	dex.routerv2, err = uniswapv2.NewIUniswapV2Router02(common.Address(common.HexToAddress(uniswapv2RouterAddress)), client)
	if (err != nil) { 
		fmt.Printf("Failed to get router at %v",uniswapv2RouterAddress)
		return err
	}
	return nil
}
	

func (dex *DexUniswapV2)GetSwapPools(ctx *DefaultApiService) (restapi.ImplResponse, error) {
	client := ctx.client
	swapPools := make([]restapi.SwapPool, 0)
	for i, token0 :=  range uniswapV2TokenList {
		for j, token1 :=  range uniswapV2TokenList[i+1:] {
			if (token0 != token1) {
				fmt.Printf("Pair for %v %v\n", token0.name, token1.name)
				opts  := bind.CallOpts{Pending: false}
				pairAddr, err := dex.factory.GetPair(&opts, token0.address, token1.address)
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
						Id: int64(i*1000 + j), 
						Name: pairAddr.Hex(),
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
	return restapi.Response(http.StatusOK, swapPools), nil	
}