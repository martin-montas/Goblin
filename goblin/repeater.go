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
	url.TitleStyle.Fg = ui.ColorYellow
	return url
}


