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

	// drawText(screen, 0, 1, 80, 1, defStyle, "Sup Alec")

	testBox := Box {
		startX: 20,
		startY: 5,
		endX: 60,
		endY: 25,
		style: defStyle,
	}

        box2 := MessageBox {
		Box: Box{startX: 20,
		startY: 26,
		endX: 60,
		endY: 30,
		style: defStyle,
		},
	}

	box2.messages = append(box2.messages, "Test Message")

	drawBox(screen, testBox)
	drawBox(screen, box2.Box)
        drawText(screen, box2.startX + 2, box2.startY +1, box2.endX, box2.endY, box2.style, box2.messages[0])

	screen.Show()

	time.Sleep(3 * time.Second)
	screen.Fini()
}
