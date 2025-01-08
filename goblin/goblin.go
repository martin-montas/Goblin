package goblin

import (
	"github.com/atotto/clipboard"
	ui "github.com/gizak/termui/v3"
	"log"
)

func AppRender() {
	uiEvents := ui.PollEvents()
	var inputBuffer string
	urlWidget := URLWidget()
	tabWidget := tabWidget()
	requestBox := requestBox()
	responseBox := responseBox()

	ui.Render(
		urlWidget, tabWidget,
		responseBox, requestBox,

		)

	render_tab := func() {
		switch tabWidget.ActiveTabIndex {
		case 0:
			ui.Render(urlWidget)
		case 1:
			// ui.Render(bc)
	
		case 2:
			// ui.Render(bc)
		}
	}

	for {
		e := <-uiEvents
		switch e.Type {
		case ui.KeyboardEvent:
			switch e.ID {
			case "<C-c>": // Exit on Ctrl+C
				return

			case "<Enter>": // Handle Enter key
				urlWidget.Text = inputBuffer
				ui.Render(urlWidget)

			case "<Backspace>": // Handle backspace
				if len(inputBuffer) > 0 {
					inputBuffer = inputBuffer[:len(inputBuffer)-1]
				}

			case "<C-h>", "<Left>": // Handle tab (left)
				tabWidget.FocusLeft()
				ui.Clear()
				// TODO
				// ui.Render(urlWidget,tabWidget)
				render_tab()

			case "<C-l>", "<Right>": // Handle tab (right)
				tabWidget.FocusRight()
				ui.Clear()
				render_tab()

			case "?": // help

			case "<C-p>": // paste
				clipboardContent, err := clipboard.ReadAll()
				if err == nil {
					inputBuffer += clipboardContent
				} else {
					log.Fatal(err)
				}

			default:
				// Append typed character
				inputBuffer += e.ID
			}

			// Update the text box dynamically
			urlWidget.Text = inputBuffer
			ui.Render(urlWidget)
		}
	}
}
