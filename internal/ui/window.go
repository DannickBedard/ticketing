package ui

import (
    "fmt"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"

    "github.com/DannickBedard/ticketing/config"
    "github.com/DannickBedard/ticketing/internal/browser"
)

func OpenInputWindow() {
    a := app.New()
    w := a.NewWindow("Open Ticket")

    entry := widget.NewEntry()
    entry.SetPlaceHolder("Enter ticket number")

    button := widget.NewButton("Open", func() {
        ticket := entry.Text
        url := fmt.Sprintf("%s/%s", config.BaseURL, ticket)
        browser.Open(url)
        w.Close()
    })

    w.SetContent(container.NewVBox(entry, button))
    w.Resize(fyne.NewSize(300, 100))
    w.ShowAndRun()
}

