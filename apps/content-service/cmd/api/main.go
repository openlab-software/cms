package main

import (
	"fmt"

	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/router"
)

func main() {
	err := router.Init()
	if err != nil {
		fmt.Println(err)
	}
}
