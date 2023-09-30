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
  
  drawText(screen, 0 ,1 ,80 , 1, defStyle, "Yooooooo")  

  drawBox(screen, 20, 5, 40, 15, defStyle, "I said yoooo")

  screen.Show()

  time.Sleep(3 * time.Second)
  screen.Fini()
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
  row := y1
  col := x1
  for _, r := range []rune(text) {
    s.SetContent(col, row, r, nil, style)
    col++
    if col >= x2 {
      row++
      col = x1
    }
    if row > y2 {
      break
    }
  }
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}
