

package dexes

import (
    "log"
	"time"
    "github.com/kelseyhightower/envconfig"	
)
 
type Config struct {
	LogLevel	string `default:"trace"`
    Timeout     time.Duration `default:"10m"`
    InfuraId  string `required:"true"`
}

func NewConfig() Config {
	var config Config
    err := envconfig.Process("dex", &config)	
    if err != nil {
        log.Fatal(err.Error())
    }
	return config
}