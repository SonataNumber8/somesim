package main

import (
	"log"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	msgBoxStyle := tcell.StyleDefault.Background(tcell.ColorGray).Foreground(tcell.ColorSnow)

	screen.Init()
	screen.Clear()

	testBox := Box {
		startX: 1,
		startY: 0,
		endX: 60,
		endY: 25,
		style: defStyle,
	}

        box2 := MessageBox {
		Box: Box{startX: 1,
		startY: 26,
		endX: 60,
		endY: 30,
		style: msgBoxStyle,
		},
	}

	box2.messages = append(box2.messages, "Test Message")

	drawBox(screen, testBox)
	drawBox(screen, box2.Box)
        drawText(screen, box2.startX + 2, box2.startY +1, box2.endX, box2.endY, box2.style, box2.messages[0])

	screen.Show()
        for {
		screen.Show()
		event := screen.PollEvent()
		
		switch event := event.(type) {
			case *tcell.EventKey:
				if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
					exit(screen)
					fmt.Println("Exiting")
					return
				}
			case *tcell.EventResize:
				screen.Sync()
		}
	}

}

func exit(screen tcell.Screen) {
	screen.Fini()
}
