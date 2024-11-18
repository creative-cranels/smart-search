package server

import (
	"fmt"

	"github.com/creative-cranels/smart-search/config"
)

func Start(srvConfig ServerConfig, cfg *config.Config) {
	fmt.Println("start with srv config:", srvConfig)
	fmt.Printf("start with config: %+v\n", cfg)
}
