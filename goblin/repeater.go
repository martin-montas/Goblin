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


func requestBox() *widget.Paragraph {
	width, height := ui.TerminalDimensions()
	requestBox := widget.NewParagraph()
	requestBox.Title = "Request"

	requestBox.SetRect(0, 8, (width/2) -2, height -5)
	requestBox.TitleStyle.Fg = ui.ColorRed

	return requestBox
}

func responseBox() *widget.Paragraph {
	width, height := ui.TerminalDimensions()
	responseBox := widget.NewParagraph()
	responseBox.Title = "Response"

	responseBox.SetRect(60, 8, width, height -5)
	responseBox.TitleStyle.Fg = ui.ColorRed

	return responseBox
}



