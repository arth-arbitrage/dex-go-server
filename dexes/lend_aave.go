package dexes

import (
	"net/http"
	"fmt"
	"math"
	//"errors"
	//"context"
	"math/big"	
	//"encoding/json"
	"github.com/arth-arbitrage/dex-go-server/go-contracts/aave"
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"	
	"github.com/arth-arbitrage/dex-go-server/go"
)

type LendAave struct {
	name string
	addressProvider *aave.ILendingPoolAddressesProvider
	lendingPool *aave.ILendingPool
}

const aaveAddressProviderAddr = "0x24a42fD28C976A61Df5D00D0599C34c4f90748c8"

func (lend *LendAave) Name(ctx *DefaultApiService) string {
	return lend.name
}


func (lend *LendAave)Init(ctx *DefaultApiService) error {
	var err error
	client:= ctx.client
	opts  := bind.CallOpts{Pending: false}
	lend.addressProvider, err = aave.NewILendingPoolAddressesProvider(common.Address(common.HexToAddress(aaveAddressProviderAddr)), client)
	if (err != nil) { 
		fmt.Printf("Failed to get aave address provider at %v",aaveAddressProviderAddr)
		return  err 
	}
	lendingPoolAddr, err := lend.addressProvider.GetLendingPool(&opts)
	if (err != nil) { 
		fmt.Printf("Failed to get lendingPoolAddr")
		return  err 
	}		

	lend.lendingPool, err = aave.NewILendingPool(lendingPoolAddr, client)
	if (err != nil) { 
		fmt.Printf("Failed to create lendingPool at %v",lendingPoolAddr)
		return  err 
	}		
	return nil
}


func (lend *LendAave)GetLenderPools(ctx *DefaultApiService) (restapi.ImplResponse, error) {
	lenderPools := make([]restapi.LenderPool, 0)
	for i, token :=  range aaveTokenList {

		fmt.Printf("Lending pool for %v\n", token.name,)
		opts  := bind.CallOpts{Pending: false}

		reserves, err := lend.lendingPool.GetReserveData(&opts, token.address)
		if (err != nil) { 
			fmt.Printf("Failed to GetReserveData for token%v", token.address)
			continue
		} 

		t0_div := new(big.Float).SetFloat64(math.Pow(10, float64(token.decimals)))

		f0 := new(big.Float).SetInt(reserves.AvailableLiquidity)
		f0 = new(big.Float).Quo(f0, t0_div)
		f0_f32, _ := f0.Float32()

		rate := new(big.Float).SetInt(reserves.FixedBorrowRate)
		rate = new(big.Float).Quo(rate, t0_div)
		rate_f32, _ := rate.Float32()

		lenderPool := restapi.LenderPool{
			Id: int64(i), 
			Name: token.address.Hex(),
			Token: token.name,
			Reserve: f0_f32,
			Fees: rate_f32,
		}

		lenderPools = append(lenderPools, lenderPool)
			
		
	}

	return restapi.Response(http.StatusOK, lenderPools), nil	
}