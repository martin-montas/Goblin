package main

import (
	ui "github.com/gizak/termui/v3"
	goblin "goblin/goblin"
	"log"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	goblin.AppRender()
}
