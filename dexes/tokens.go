
package dexes

import (
	"github.com/ethereum/go-ethereum/common"
)

type token struct {
	name string
	address  common.Address
	decimals int
}



var uniswapV2TokenList = []token {
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

var curveTokenList = []token {
	//token {name: "WETH", address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), decimals:18 },
    token {name: "DAI", address:common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"), decimals:18},
    token {name: "USDC", address: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"), decimals:6},
    token {name: "USDT", address: common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"), decimals:6},
    token {name: "wBTC", address: common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"), decimals:8},
    token {name: "renBTC", address: common.HexToAddress("0xeb4c2781e4eba804ce9a9803c67d0893436bb27d"), decimals:8},
    token {name: "BUSD", address: common.HexToAddress("0x4fabb145d64652a948d72533023f6e7a623c7c53"), decimals:18},
    token {name: "TUSD", address: common.HexToAddress("0x0000000000085d4780b73119b644ae5ecd22b376"), decimals:18},
    token {name: "EURS", address: common.HexToAddress("0xdB25f211AB05b1c97D595516F45794528a807ad8"), decimals:2},
    //token {name: "VID", address: common.HexToAddress("0x2c9023bbc572ff8dc1228c7858a280046ea8c9e5"), decimals:18},
    token {name: "aETH", address: common.HexToAddress("0xE95A203B1a91a908F9B9CE46459d101078c2c3cb"),  decimals:18},    
    token {name: "ETH", address: common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),  decimals:18},
}
