package main

import "fmt"
import "os"
import "github.com/nsf/termbox-go"
import "time"
import "net"

func panicIfError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  port := "10001"
  if len(os.Args) > 1 {
    port = os.Args[1]
  }

  peerAddr := "127.0.0.1:10002"
  if len(os.Args) > 2 {
    peerAddr = os.Args[2]
  }

  err := termbox.Init()
  panicIfError(err)
  defer termbox.Close()

  addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%s", port))
  panicIfError(err)

  conn, err := net.ListenUDP("udp", addr)
  panicIfError(err)
  defer conn.Close()

  netQueue := make(chan []byte)
  go func() {
    buf := make([]byte, 1024)
    for {
      n, _, err := conn.ReadFromUDP(buf)
      if err == nil {
        netQueue <- buf[0:n]
      }
    }
  }()

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
          endGame()
          startGame()
        case 'p':
          peer, err := net.ResolveUDPAddr("udp", peerAddr)
          local, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
          if err == nil {
            peerConn, err := net.DialUDP("udp", local, peer)
            if err == nil {
              buf := []byte("Hello")
              _, err = peerConn.Write(buf)
              if err != nil {
                setScore(err.Error())
              }
              peerConn.Close()
            } else {
              setScore(err.Error())
            }
          } else {
            setScore(err.Error())
          }
        }
      }
    case buf := <- netQueue:
      setScore(string(buf))
    default:
      termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
      drawGame()
      termbox.Flush()

      time.Sleep(10 * time.Millisecond)
    }
  }
}
