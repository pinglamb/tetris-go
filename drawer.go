package main

import "github.com/nsf/termbox-go"

const cellWidth = 5
const cellHeight = 2

const boardWidth = 10
const boardHeight = 25

func drawTetromino(t Tetromino, x, y int) {
  color := TetrominoColors[t]
  for _, point := range TetrominoShapes[t] {
    drawPoint(x + point[0], y + point[1], color)
  }
}

func drawPoint(x, y int, color termbox.Attribute) {
  for i := 0; i < cellWidth; i++ {
    for j := 0; j < cellHeight; j++ {
      termbox.SetCell(x * cellWidth + i, y * cellHeight + j, ' ', color, color)
    }
  }
}
