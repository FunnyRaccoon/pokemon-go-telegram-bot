package main

import (
	"fmt"

	"github.com/mjibson/gpsoauth"

	"github.com/FunnyRaccoon/pokemon-go-telegram-bot/app"
	"github.com/FunnyRaccoon/pokemon-go-telegram-bot/app/config"
)

func main() {
	c, err := config.ReadConfig("config_dev.json")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%#v", c)
	auth, err := Login(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(auth)
	fmt.Println(app.Resources)
}

func Login(c *config.Config) (string, error) {
	return gpsoauth.Login(
		c.UserConfig.Username,
		c.UserConfig.Password,
		c.AppConfig.AndroidId,
		c.AppConfig.Service,
	)
}
