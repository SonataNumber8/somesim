package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	//defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	msgBoxStyle := tcell.StyleDefault.Background(tcell.ColorGray).Foreground(tcell.ColorSnow)

	screen.Init()
	screen.Clear()

	box2 := MessageBox{
		Box: Box{startX: 1,
			startY: 1,
			endX:   90,
			endY:   10,
			style:  msgBoxStyle,
		},
	}

	drawBox(screen, box2.Box)

	screen.Show()
	i := 1
	for {
		box2.messages = append(box2.messages, "test message" + strconv.Itoa(i))
		i++
		screen.Show()
		event := screen.PollEvent()

		switch event := event.(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				fmt.Println("Exiting")
				return
			}
		case *tcell.EventResize:
			drawBox(screen, box2.Box)
			box2.ShowMessages(screen)
			screen.Sync()
		}
	}
}
