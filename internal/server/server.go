package server

import (
	"fmt"

	"github.com/creative-cranels/smart-search/configs"
)

func Start(srvConfig ServerConfig, cfg *configs.Config) {
	fmt.Println("start with srv config:", srvConfig)
	fmt.Printf("start with config: %+v\n", cfg)
}
