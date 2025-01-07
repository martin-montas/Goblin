package main

import (
	"log"
	ui "github.com/gizak/termui/v3"
	g "goblin/goblin"
)


func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	g.AppRender()
}








