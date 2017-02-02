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

  for {
    select {
    case e := <- eventQueue:
      if e.Type == termbox.EventKey && (e.Ch == 'q' || e.Key == termbox.KeyEsc) {
        return
      }
    default:
      termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

      drawPanes()

      drawTetrominoOnBoard(TetrominoI, 0, 4, 0)
      drawTetrominoOnNextPane(TetrominoI)
      drawTetrominoOnHoldPane(TetrominoT)

      drawBoard()

      termbox.Flush()

      time.Sleep(10 * time.Millisecond)
    }
  }
}
