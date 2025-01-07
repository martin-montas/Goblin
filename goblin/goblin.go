package goblin

import (
	ui "github.com/gizak/termui/v3"
	"github.com/atotto/clipboard"
	"log"
)

func AppRender() {
	uiEvents := ui.PollEvents()
	var inputBuffer string

	urlWidget := URLWidget()
	ui.Render(urlWidget)

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
			case "<Tab>": // Handle tab

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

