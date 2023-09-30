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
  
  screen.SetContent(0, 20, 'H', nil, defStyle)
  screen.SetContent(1, 20, 'i', nil, defStyle)

  screen.Show()

  time.Sleep(3 * time.Second)
  screen.Fini()
}
