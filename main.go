package main

import "flag"
import "github.com/nsf/termbox-go"
import "time"

func panicIfError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  flag.StringVar(&myPort, "port", "10001", "Port to bind as host")
  flag.StringVar(&peerInfo, "peer", "", "Peer IP/Port to connect")
  flag.Parse()

  asPeer = peerInfo != ""

  err := termbox.Init()
  panicIfError(err)
  defer termbox.Close()

  initHost()
  defer closeHost()

  connectPeerIfSet()
  pingPeer()
  defer closePeer()

  eventQueue := make(chan termbox.Event)
  go func() {
    for {
      eventQueue <- termbox.PollEvent()
    }
  }()

  for {
    select {
    case e := <- eventQueue:
      if e.Type == termbox.EventKey {
        if (e.Ch == 'q' || e.Key == termbox.KeyEsc) {
          endGame()
          return
        }
      }

      if !dead {
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
        case 'n':
          if !asPeer {
            endGame()
            startGame()
          }
        }
      }
    default:
      termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
      drawGame()
      termbox.Flush()

      time.Sleep(10 * time.Millisecond)
    }
  }
}
