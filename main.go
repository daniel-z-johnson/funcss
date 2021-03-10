package main

import (
	"fmt"
	"github.com/daniel-z-johnson/funcss/models"
)

func main() {
	fmt.Println("Starting funcss")
	badAss, _ := models.FirstBadAssPage()
	fmt.Println(badAss)
}
