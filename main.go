package main

import (
	"flag"
	"fmt"
	"github.com/daniel-z-johnson/funcss/conf"
	"github.com/daniel-z-johnson/funcss/models"
)

func main() {
	confFile := flag.String("conf", "config.json", "The location of the config file")
	conf, err := conf.LoadFunCSSConf(*confFile)
	if err != nil {
		fmt.Printf("Error: %s", err)
		panic(err)
	}
	fmt.Println(conf)
	services := models.NewServices(*confFile)
	fmt.Println("Starting funCSS")
	funCESSES, err := models.GetAllBadAssAsFunCSS()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(funCESSES)
	for _, fcss := range funCESSES {
		services.FunCSS.Create(fcss)
	}
	fmt.Println(services.FunCSS.ByID(129))
}
