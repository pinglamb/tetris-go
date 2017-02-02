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

var Tetrominos = []Tetromino {
  TetrominoI,
  TetrominoJ,
  TetrominoL,
  TetrominoO,
  TetrominoS,
  TetrominoT,
  TetrominoZ,
}

var TetrominoShapes = [][][][]int {
  {{{1}, {1}, {1}, {1}}, {{1, 1, 1, 1}}},
  {{{0, 1}, {0, 1}, {1, 1}}, {{1}, {1, 1, 1}}, {{1, 1}, {1}, {1}}, {{1, 1, 1}, {0, 0, 1}}},
  {{{1}, {1}, {1, 1}}, {{1, 1, 1}, {1}}, {{1, 1}, {0, 1}, {0, 1}}, {{0, 0, 1}, {1, 1, 1}}},
  {{{1, 1}, {1, 1}}},
  {{{1}, {1, 1}, {0, 1}}, {{0, 1, 1}, {1, 1}}},
  {{{0, 1}, {1, 1, 1}}, {{1}, {1, 1}, {1}}, {{1, 1, 1}, {0, 1}}, {{0, 1}, {1, 1}, {0, 1}}},
  {{{0, 1}, {1, 1}, {1}}, {{1, 1}, {0, 1, 1}}},
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
