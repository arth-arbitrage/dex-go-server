package dexes

import (
	"net/http"
	"os"
	"io/ioutil"	
	//"fmt"
	//"math"
	//"errors"
	//"context"
	"math/big"	
	//"encoding/json"
	"github.com/arth-arbitrage/dex-go-server/go-contracts/arbitrage"	
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"	
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/arth-arbitrage/dex-go-server/go"
)

func transactor(keyfile, password string) (*bind.TransactOpts, error) {
	f, err := os.Open(keyfile)
	if err != nil {
			return nil, err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
			return nil, err
	}
	key, err := keystore.DecryptKey(data, password)
	if err != nil {
			return nil, err
	}
	return bind.NewKeyedTransactor(key.PrivateKey), nil
}

func Multiswap(ctx *DefaultApiService, multiSwap restapi.MultiSwapExec) (restapi.ImplResponse, error) {
	client := ctx.client
	opts, err   := transactor("testfile", "testkey") //*bind.TransactOpts
	if err != nil {
		return restapi.Response(http.StatusBadRequest, err), err
	}	
	path := make([]common.Address, 0)
	amounts := make([]*big.Int, 0)
	loanWrapAddress := common.HexToAddress(multiSwap.Lenderwrap)
	lenderAddress := common.HexToAddress(multiSwap.Lender)
	arbContract, err := arbitrage.NewIArthArbV1MultiSwap(loanWrapAddress, client)
	if err != nil {
		return restapi.Response(http.StatusBadRequest, err), err
	}	

	amountIn := Float32Coin( multiSwap.Amount, int(multiSwap.Decimals))

	for _, swap := range multiSwap.Swaps {
		amount := Float32Coin(swap.Amount, int(swap.Decimals1))
		amounts = append(amounts, amount)
		asset := common.HexToAddress(swap.Token1Address)
		path = append(path,asset)
	}

	loanAsset := path[0]
	txn, err := arbContract.Arbitrage(opts, lenderAddress, loanAsset, amountIn, path, amounts)
	if err != nil {
		return restapi.Response(http.StatusBadRequest, err), err
	}		
	return restapi.Response(http.StatusOK, txn.Hash().Hex()), nil	
}