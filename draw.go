package main

import "github.com/nsf/termbox-go"

const cellWidth = 5
const cellHeight = 2

const boardWidth = 10
const boardHeight = 25

var board [boardWidth][boardHeight]termbox.Attribute

func drawTetromino(t Tetromino, x, y, spin int) {
  color := TetrominoColors[t]
  normalizedSpin := spin % len(TetrominoShapes[t])
  for r, cols := range TetrominoShapes[t][normalizedSpin] {
    for c, flag := range cols {
      if flag == 1 {
        drawPoint(x + c, y + r, color)
      }
    }
  }
}

func drawBoard() {
  for i := 0; i < boardWidth; i++ {
    for j := 0; j < boardHeight; j++ {
      board[i][j] = termbox.ColorDefault
    }
  }

  board[0][21] = termbox.ColorCyan
  board[0][22] = termbox.ColorCyan
  board[0][23] = termbox.ColorCyan
  board[0][24] = termbox.ColorCyan
  board[1][22] = termbox.ColorRed
  board[1][23] = termbox.ColorRed
  board[1][24] = termbox.ColorRed
  board[2][24] = termbox.ColorRed

  for i := 0; i < boardWidth; i++ {
    for j := 0; j < boardHeight; j++ {
      drawPoint(i, j, board[i][j])
    }
  }
}

func drawPoint(x, y int, color termbox.Attribute) {
  for i := 0; i < cellWidth; i++ {
    for j := 0; j < cellHeight; j++ {
      termbox.SetCell(x * cellWidth + i, y * cellHeight + j, ' ', color, color)
    }
  }
}
