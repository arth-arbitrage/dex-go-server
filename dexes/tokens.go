
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

/*
("0x6B175474E89094C44Da98b954EedeAC495271d0F", 
"0x0000000000085d4780B73119b644AE5ecd22b376", 
"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", 
"0xdAC17F958D2ee523a2206206994597C13D831ec7", 
"0x57Ab1ec28D129707052df4dF418D58a2D46d5f51", 
"0x80fB784B7eD66730e8b1DBd9820aFD29931aab03", 
"0x0D8775F648430679A709E98d2b0Cb6250d2887EF", 
"0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", 
"0x514910771AF9Ca656af840dff83E8264EcF986CA", 
"0xdd974D5C2e2928deA5F71b9825b8b646686BD200", 
"0x1985365e9f78359a9B6AD760e32412f4a445E862", 
"0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2", 
"0x0F5D2fB29fb7d3CFeE444a200298f468908cC942", 
"0xE41d2489571d322189246DaFA5ebDe1F4699F498", 
"0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F", 
"0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599", 
"0x4Fabb145d64652a948d72533023f6E7A623C7C53", 
"0xF629cBd94d3791C9250152BD8dfBDF380E2a3B9c", 
"0x408e41876cCCDC0F92210600ef50372656052a38", 
"0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e", 
"0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9", 
"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")
*/
var aaveTokenList = []token {
    token {name: "DAI", address:common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"), decimals:18},
    token {name: "USDC", address: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"), decimals:6},
    token {name: "USDT", address: common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"), decimals:6},
    token {name: "wBTC", address: common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"), decimals:8},
    token {name: "BUSD", address: common.HexToAddress("0x4fabb145d64652a948d72533023f6e7a623c7c53"), decimals:18},
}

var arthTokenList = []token {
    token {name: "DAI", address:common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"), decimals:18},
    token {name: "USDC", address: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"), decimals:6},
    token {name: "USDT", address: common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"), decimals:6},
    token {name: "wBTC", address: common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"), decimals:8},
    token {name: "BUSD", address: common.HexToAddress("0x4fabb145d64652a948d72533023f6e7a623c7c53"), decimals:18},
}