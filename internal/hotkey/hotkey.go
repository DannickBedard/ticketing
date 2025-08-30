package hotkey

import (
    "log"

    hook "github.com/robotn/gohook"
    "github.com/DannickBedard/ticketing/internal/ui"
)

func Listen() {
    evChan := hook.Start()
    defer hook.End()

    for ev := range evChan {
        // ctrl + alt + t
        if ev.Kind == hook.KeyDown &&
            ev.Rawcode == 0x54 && // T key
            ev.Ctrl && ev.Alt {
            log.Println("Hotkey pressed: opening ticket window")
            go ui.OpenInputWindow()
        }
    }
}

