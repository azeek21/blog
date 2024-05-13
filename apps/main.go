package main

import (
	"fmt"

	"github.com/azeek21/blog/apps/drawer"
)

func main() {
	err := drawer.DrawPngFromFile("./drawer/image-example.dws", "")
	if err != nil {
		fmt.Println(err.Error())
	}
}
