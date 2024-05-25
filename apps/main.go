package main

import (
	"fmt"

	"github.com/azeek21/blog/apps/drawer"
)

func main() {
	drw := drawer.NewDrawer()
	err := drw.DrawPngFromFile("./logo.dws", "out.png")
	if err != nil {
		fmt.Println(err.Error())
	}
}
