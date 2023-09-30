package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	screen.Init()

	screen.Clear()

	drawText(screen, 0, 1, 80, 1, defStyle, "Sup Alec")

	//drawBox(screen, 20, 5, 60, 25, defStyle, "I made this box just now")

	testBox := Box{
		startX: 20,
		startY: 5,
		endX: 60,
		endY: 25,
		style: defStyle,
		text: "I made this box just now", 
	}

	drawBox(screen, testBox)

	screen.Show()

	time.Sleep(3 * time.Second)
	screen.Fini()
}
