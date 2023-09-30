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
