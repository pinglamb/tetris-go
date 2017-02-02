package main

// import "fmt"
import "github.com/nsf/termbox-go"
import "time"

func main() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()

  eventQueue := make(chan termbox.Event)
  go func() {
    for {
      eventQueue <- termbox.PollEvent()
    }
  }()

  startGame()

  for {
    select {
    case e := <- eventQueue:
      if e.Type == termbox.EventKey {
        if (e.Ch == 'q' || e.Key == termbox.KeyEsc) {
          endGame()
          return
        }
      }

      switch e.Ch {
      case 'w':
        dropTetromino()
      case 's':
        moveTetrominoDown()
      case 'a':
        moveTetrominoLeft()
      case 'd':
        moveTetrominoRight()
      case 'j':
        spinTetromino()
      case 'h':
        holdTetronmino()
      }
    default:
      termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
      drawGame()
      termbox.Flush()

      time.Sleep(10 * time.Millisecond)
    }
  }
}
