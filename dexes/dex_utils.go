package dexes

import (
	"math"
	"fmt"
	"strings"
	"math/big"
	"github.com/ethereum/go-ethereum/params"	
)

func etherFloatToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func Float32Coin(value float32, decimals int) *big.Int {
	coinBigLoat := new (big.Float).SetFloat64(float64(value))
	truncInt, _ := coinBigLoat.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(int64(math.Pow10(decimals))))
	// TODO
	fracStr := strings.Split(fmt.Sprintf("%.18f", coinBigLoat), ".")[1]
	fracStr += strings.Repeat("0", decimals-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	Coin := new(big.Int).Add(truncInt, fracInt)
	return Coin
}