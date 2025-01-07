package goblin

import (
	ui "github.com/gizak/termui/v3"
	widget "github.com/gizak/termui/v3/widgets"
)

func URLWidget() *widget.Paragraph {
	width, _ := ui.TerminalDimensions()
	url := widget.NewParagraph()
	url.Title = "Url"
	url.SetRect(0, 3, width, 6)
	url.TitleStyle.Fg = ui.ColorRed

	return url
}

func tabWidget() *widget.TabPane {
	width, _ := ui.TerminalDimensions()
	tabpane := widget.NewTabPane("Repeater", "Intruder", "About")
	tabpane.SetRect(0, 0, width, 3)
	tabpane.ActiveTabIndex = 0
	tabpane.TitleStyle.Fg = ui.ColorRed
	tabpane.Border = false

	return tabpane
}
