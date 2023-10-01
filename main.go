package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	screen2, _ := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	//defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	msgBoxStyle := tcell.StyleDefault.Background(tcell.ColorGray).Foreground(tcell.ColorSnow)

	screen.Init()
	screen.Clear()
	screen2.Init()
	screen2.Clear()

	box2 := MessageBox{
		Box: Box{startX: 1,
			startY: 1,
			endX:   90,
			endY:   10,
			style:  msgBoxStyle,
		},
	}

	box3 := Box{
		startX: 1,
		startY: 1,
		endX:   90,
		endY:   10,
		style:  msgBoxStyle,
		text: "Screen 2",
	}

	drawBox(screen2, box3)

	drawBox(screen, box2.Box)

	screen.Show()
	i := 1
	for {
		drawBox(screen, box2.Box)
                box2.ShowMessages(screen)
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
			} else if event.Rune() == 'C' || event.Rune() == 'c' {
				drawBox(screen2, box3)
				screen2.Sync()
				screen2.Show()
				time.Sleep(5 * time.Second)
				screen.Sync()
			}
		case *tcell.EventResize:
			screen.Sync()
		}
	}
}
