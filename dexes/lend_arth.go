package dexes

import (
	"net/http"
	"fmt"
	"math"
	//"errors"
	//"context"
	"math/big"	
	//"encoding/json"
	"github.com/arth-arbitrage/dex-go-server/go-contracts/arth"
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"	
	"github.com/arth-arbitrage/dex-go-server/go"
)

type LendArth struct {
	name string
	lendingPool *arth.ArthLending
}

const arthLenderWrapAddress = "0x24a42fD28C976A61Df5D00D0599C34c4f90748c8"
const lenderAddress = "0x24a42fD28C976A61Df5D00D0599C34c4f90748c8"

func (lend *LendArth) Name(ctx *DefaultApiService) string {
	return lend.name
}


func (lend *LendArth)Init(ctx *DefaultApiService) error {
	var err error
	client:= ctx.client
	//opts  := bind.CallOpts{Pending: false}
		

	lend.lendingPool, err = arth.NewArthLending(common.Address(common.HexToAddress(arthLenderWrapAddress)), client)
	if (err != nil) { 
		fmt.Printf("Failed to create lendingPool at %v",arthLenderWrapAddress)
		return  err 
	}		
	return nil
}


func (lend *LendArth)GetLenderPools(ctx *DefaultApiService) (restapi.ImplResponse, error) {
	lenderPools := make([]restapi.LenderPool, 0)
	for i, token :=  range arthTokenList {

		fmt.Printf("Lending pool for %v\n", token.name,)
		opts  := bind.CallOpts{Pending: false}

		reserves, err := lend.lendingPool.GetReserveAvailableLiquidity(&opts, token.address)
		if (err != nil) { 
			fmt.Printf("Failed to GetReserveData for token%v", token.address)
			continue
		} 

		t0_div := new(big.Float).SetFloat64(math.Pow(10, float64(token.decimals)))

		f0 := new(big.Float).SetInt(reserves)
		f0 = new(big.Float).Quo(f0, t0_div)
		f0_f32, _ := f0.Float32()

		//rate := new(big.Float).SetInt(reserves.FixedBorrowRate)
		rate := new(big.Float).SetInt64(35)
		rate = new(big.Float).Quo(rate, t0_div)
		rate_f32, _ := rate.Float32()

		lenderPool := restapi.LenderPool{
			Id: int64(i), 
			Name: token.address.Hex(),
			Token: token.name,
			TokenAddress: token.address.Hex(),
			LenderWrap: arthLenderWrapAddress,
			Lender: lenderAddress,
			Reserve: f0_f32,
			Fees: rate_f32,
		}

		lenderPools = append(lenderPools, lenderPool)	
	}
	return restapi.Response(http.StatusOK, lenderPools), nil	
}