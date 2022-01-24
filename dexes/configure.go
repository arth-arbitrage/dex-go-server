

package dexes

import (
    "log"
	"time"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "github.com/kelseyhightower/envconfig"	
)
 
type Config struct {
	LogLevel	string `default:"trace"`
    Timeout     time.Duration `default:"10m"`
    InfuraId    string `required:"true"`
    // Contract Addresses
    AaveAddressProvider string  `default:"0x24a42fD28C976A61Df5D00D0599C34c4f90748c8"`
    AaveArbMultiSwapV1  string

    UniswapV2Factory    string  `default:"0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"`
    UniswapV2Router02   string  `default:"0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"`
    Uniswapv2SwapWrap   string

    CurveAddressProvider string `default:"0x0000000022D53366457F9d5E68Ec105046FC4383"`
    CurveSwapWrap        string

    ArthArbV1MultiSwap  string
    ArthLending         string
}

// get configuration from environment variable
func NewConfig() Config {
	var config Config
    err := envconfig.Process("dex", &config)	
    if err != nil {
        log.Fatal(err.Error())
    }
	return config
}

func GetConfig(path string) (*Config, error) {
	conf := Config{}
	if len(path) > 0 {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("can't read config at %s: %w", path, err)
		}
		if err := json.Unmarshal(data, &conf); err != nil {
			return nil, err
		}
	}
	if err := envconfig.Process("dex", &conf); err != nil {
		return nil, fmt.Errorf("failed to parse envconfig: %w", err)
	}
    return &conf, nil
}