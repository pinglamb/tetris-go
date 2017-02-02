package main

import "github.com/nsf/termbox-go"

type Tetromino int

const (
  TetrominoI Tetromino = iota
  TetrominoJ
  TetrominoL
  TetrominoO
  TetrominoS
  TetrominoT
  TetrominoZ
)

var TetrominoShapes = [][][]int {
  {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
  {{1, 0}, {1, 1}, {1, 2}, {0, 2}},
  {{0, 0}, {0, 1}, {0, 2}, {1, 2}},
  {{0, 0}, {0, 1}, {1, 0}, {1, 1}},
  {{0, 0}, {0, 1}, {1, 1}, {1, 2}},
  {{1, 0}, {0, 1}, {1, 1}, {2, 1}},
  {{0, 1}, {1, 0}, {1, 1}, {0, 2}},
}

var TetrominoColors = []termbox.Attribute {
  termbox.ColorCyan,
  termbox.ColorBlue,
  termbox.ColorYellow,
  termbox.ColorWhite,
  termbox.ColorGreen,
  termbox.ColorRed,
  termbox.ColorMagenta,
}
