package main

import (
	"github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, startX int, startY int, endX int, endY int, style tcell.Style, text string) {
	row := startY
	col := startX
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= endX {
			row++
			col = startX
		}
		if row > endY {
			break
		}
	}
}

func drawBox(s tcell.Screen, box Box) {

	// Fill background
	for row := box.startY; row <= box.endY; row++ {
		for col := box.startX; col <= box.endX; col++ {
			s.SetContent(col, row, ' ', nil, box.style)
		}
	}

	// Draw borders
	for col := box.startX; col <= box.endX; col++ {
		s.SetContent(col, box.startY, tcell.RuneHLine, nil, box.style)
		s.SetContent(col, box.endY, tcell.RuneHLine, nil, box.style)
	}
	for row := box.startY + 1; row < box.endY; row++ {
		s.SetContent(box.startX, row, tcell.RuneVLine, nil, box.style)
		s.SetContent(box.endX, row, tcell.RuneVLine, nil, box.style)
	}

	// Only draw corners if necessary
	if box.startY != box.endY && box.startX != box.endX {
		s.SetContent(box.startX, box.startY, tcell.RuneULCorner, nil, box.style)
		s.SetContent(box.endX, box.startY, tcell.RuneURCorner, nil, box.style)
		s.SetContent(box.startX, box.endY, tcell.RuneLLCorner, nil, box.style)
		s.SetContent(box.endX, box.endY, tcell.RuneLRCorner, nil, box.style)
	}

	if box.text != "" {
		drawText(s, box.startX+1, box.startY+1, box.endX-1, box.endY-1, box.style, box.text)
	}
}

type Box struct {
	startX, startY, endX, endY int
	style tcell.Style
        text string
}

type MessageBox struct {
	Box
	messages []string
}

func (msgBox MessageBox) ShowMessages() {
	
}
