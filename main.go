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
	fmt.Println("Starting funCSS")
	badAss, _ := models.FirstBadAssPage()
	fmt.Println(badAss)
}
