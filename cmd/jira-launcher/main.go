package main

import (
    "log"

    "github.com/DannickBedard/ticketing/internal/hotkey"
)

func main() {
    log.Println("Jira Launcher is running... (Ctrl+Alt+T to open)")
    hotkey.Listen()
}

