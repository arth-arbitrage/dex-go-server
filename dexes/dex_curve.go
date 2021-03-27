package dexes

import (
	"net/http"
	"fmt"
	"math"
	//"errors"
	//"context"
	"math/big"	
	//"encoding/json"
	"github.com/arth-arbitrage/dex-go-server/go-contracts/curvefi"
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"	
	"github.com/arth-arbitrage/dex-go-server/go"
)

type DexCurve struct {
	name string
	addressProvider *curvefi.AddressProvider
	registry *curvefi.Registry
	poolinfo *curvefi.PoolInfo
	swaps *curvefi.Swaps
	
}

const curveAddressProviderAddr = "0x0000000022D53366457F9d5E68Ec105046FC4383"


func NewDexCurve() *DexCurve {
	return &DexCurve{ 
		name: "curve",
	}
}

func (dex DexCurve) Name() string {
	return dex.name
}

func (dex *DexCurve)GetSwapPools(ctx *DefaultApiService) (restapi.ImplResponse, error) {
	opts  := bind.CallOpts{Pending: false}
	client := ctx.client
	var err error
	dex.addressProvider, err = curvefi.NewAddressProvider(common.Address(common.HexToAddress(curveAddressProviderAddr)), client)
	if (err != nil) { 
		fmt.Printf("Failed to get curve address provider at %v",curveAddressProviderAddr)
		return restapi.Response(http.StatusInternalServerError, nil), err 
	}
	registryAddr, err := dex.addressProvider.GetRegistry(&opts)
	if (err != nil) { 
		fmt.Printf("Failed to cregetate Registry")
		return restapi.Response(http.StatusInternalServerError, nil), err 
	}		

	dex.registry, err = curvefi.NewRegistry(registryAddr, client)
	if (err != nil) { 
		fmt.Printf("Failed to create registry at %v",registryAddr)
		return restapi.Response(http.StatusInternalServerError, nil), err 
	}
	poolinfoAddr, err:= dex.addressProvider.GetAddress(&opts, big.NewInt(1))
	if (err != nil) { 
		fmt.Printf("Failed to get pool address %v",1)
		return restapi.Response(http.StatusInternalServerError, nil), err 
	}	
	dex.poolinfo, err = curvefi.NewPoolInfo(poolinfoAddr, client)
	if (err != nil) { 
		fmt.Printf("Failed to create PoolInfo swaps address %v",poolinfoAddr)
		return restapi.Response(http.StatusInternalServerError, nil), err 
	}		

	swapsAddr, err:= dex.addressProvider.GetAddress(&opts, big.NewInt(2))
	if (err != nil) { 
		fmt.Printf("Failed to get swaps address %v",2)
		return restapi.Response(http.StatusInternalServerError, nil), err 
	}		
	dex.swaps, err = curvefi.NewSwaps(swapsAddr, client)
	if (err != nil) { 
		fmt.Printf("Failed tocreate NewSwaps address %v",swapsAddr)
		return restapi.Response(http.StatusInternalServerError, nil), err 
	}		

	//router, err := uniswapv2.NewIUniswapV2Router02(common.Address(common.HexToAddress(routerAddress)), client)
	swapPools := make([]restapi.SwapPool, 0)
	for i, token0 :=  range curveTokenList {
		for _, token1 :=  range curveTokenList[i+1:] {
			if (token0 != token1) {
				fmt.Printf("Pair for %v %v\n", token0.name, token1.name)
				opts  := bind.CallOpts{Pending: false}
				poolAddr, rate, err := dex.swaps.GetBestRate(&opts, token0.address, token1.address, big.NewInt(1))
				if (err != nil || poolAddr == common.Address(common.HexToAddress("x0"))) { 
					fmt.Printf("Failed to get pool for %v %v", token0.name, token1.name)
					continue
				}
				
				balances, err := dex.registry.GetUnderlyingBalances(&opts, poolAddr)
				if (err != nil) { 
					fmt.Printf("Failed to GetUnderlyingBalances for pool%v", poolAddr)
					continue
				} 
				coins, err := dex.registry.GetUnderlyingCoins(&opts, poolAddr)
				if (err != nil) { 
					fmt.Printf("Failed to GetUnderlyingBalances for pool%v", poolAddr)
					continue
				} 
				var reserve0 *big.Int = big.NewInt(0)
				var reserve1 *big.Int = big.NewInt(0)

				from_token, to_token := token0, token1 
				for i, coin := range coins {
					if (coin == token0.address) {
						reserve0 = balances[i]
					}
	                if (coin == token1.address) {
        	            reserve1 = balances[i]
					}
				}

				fmt.Printf("reserve0 %v reserve1 %v", reserve0, reserve1)	
				t0_div := new(big.Float).SetFloat64(math.Pow(10, float64(from_token.decimals)))
				t1_div := new(big.Float).SetFloat64(math.Pow(10, float64(to_token.decimals)))

				f0 := new(big.Float).SetInt(reserve0)
				f0 = new(big.Float).Quo(f0, t0_div)
				f0_f32, _ := f0.Float32()

				f1 := new(big.Float).SetInt(reserve1)
				f1 = new(big.Float).Quo(f1, t1_div)
				f1_f32, _ := f1.Float32()

				p0_f32 := f1_f32/f0_f32
				p1_f32 := f0_f32/f1_f32
				_ = rate

				if(reserve0.Cmp(big.NewInt(0)) == 0 || reserve1.Cmp(big.NewInt(0)) == 0) {
					continue
				}

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
	//jsonData, err := json.Marshal(swapPools)
	//if err != nil {
	//	return restapi.Response(http.StatusInternalServerError, nil), err
	//}	
	return restapi.Response(http.StatusOK, swapPools/*(string(jsonData)*/), nil	
}