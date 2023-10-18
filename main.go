package main

import (
	"fmt"

	"github.com/osmanakol/oauth2-golang/application/api"
	config "github.com/osmanakol/oauth2-golang/configuration"
)

func main() {
	config.NewConfiguration()
	fmt.Println(config.Env)

	applicationConfig := config.Env.GetRedisConfig()
	f := api.New()
	f.Listen(":" + applicationConfig.PORT)
}
