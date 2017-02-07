package main

import "time"
import "math/rand"
import "github.com/nsf/termbox-go"

var currentBoard [][boardWidth]termbox.Attribute

var currentTetromino Tetromino
var currentTetrominoSpin int
var currentTetrominoX int
var currentTetrominoY int

var nextTetromino Tetromino
var holdedTetromino Tetromino

var hasTetrominoHolded = false
var dead = false

var gravity = 1

var gameStarted = false
var gameTicker *time.Ticker

var gameScore = ""

func startGame() {
  currentBoard = [][boardWidth]termbox.Attribute {}
  for i := 0; i < boardHeight; i++ {
    currentBoard = append(currentBoard, newRow())
  }

  nextTetromino = randTetromino()
  newTetromino()

  speed := time.Duration(900 / gravity)
  gameTicker = time.NewTicker(time.Millisecond * speed)
  go func() {
    for _ = range gameTicker.C {
      tickGame()
    }
  }()

  gameStarted = true
}

func endGame() {
  if gameStarted {
    gameTicker.Stop()
    gameStarted = false
  }
}

func tickGame() {
  if isTouchingGround(currentTetrominoX, currentTetrominoY) {
    landTetromino()
    newTetromino()
  } else {
    moveTetrominoDown()
  }
}

func newTetromino() {
  addTetromino(nextTetromino)
  nextTetromino = randTetromino()
}

func addTetromino(t Tetromino) {
  currentTetromino = t
  currentTetrominoX = 5
  currentTetrominoY = 0
  currentTetrominoSpin = 0

  if !isValidMove(currentTetrominoX, currentTetrominoY, currentTetrominoSpin) {
    dead = true
    currentTetrominoY--
    for !isValidMove(currentTetrominoX, currentTetrominoY, currentTetrominoSpin) {
      currentTetrominoY--
    }
    landTetromino()
    endGame()
  }
}

func randTetromino() Tetromino {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  return Tetrominos[r.Intn(7)]
}

func holdTetronmino() {
  if hasTetrominoHolded {
    tmp := currentTetromino
    addTetromino(holdedTetromino)
    holdedTetromino = tmp
  } else {
    hasTetrominoHolded = true
    holdedTetromino = currentTetromino
    newTetromino()
  }
}

func landTetromino() {
  blocks := blocksOf(currentTetromino, currentTetrominoX, currentTetrominoY, currentTetrominoSpin)

  for _, block := range blocks {
    if block[1] >= 0 {
      currentBoard[block[1]][block[0]] = TetrominoColors[currentTetromino]
    }
  }

  clearFullRows()
}

func moveTetrominoLeft() {
  newX := currentTetrominoX - 1
  if isValidMove(newX, currentTetrominoY, currentTetrominoSpin) {
    currentTetrominoX = newX
  }
}

func moveTetrominoRight() {
  newX := currentTetrominoX + 1
  if isValidMove(newX, currentTetrominoY, currentTetrominoSpin) {
    currentTetrominoX = newX
  }
}

func moveTetrominoDown() {
  newY := currentTetrominoY + 1
  if isValidMove(currentTetrominoX, newY, currentTetrominoSpin) {
    currentTetrominoY = newY
  }
}

func dropTetromino() {
  newY := currentTetrominoY
  for !isTouchingGround(currentTetrominoX, newY) {
    newY++
  }
  currentTetrominoY = newY
  landTetromino()
  newTetromino()
}

func spinTetromino() {
  newSpin := currentTetrominoSpin + 1
  if isValidMove(currentTetrominoX, currentTetrominoY, newSpin) {
    currentTetrominoSpin = newSpin
    return
  }

  newX := currentTetrominoX - 1
  if isValidMove(newX, currentTetrominoY, newSpin) {
    currentTetrominoX = newX
    currentTetrominoSpin = newSpin
    return
  }

  newY := currentTetrominoY - 1
  if isValidMove(currentTetrominoX, newY, newSpin) {
    currentTetrominoY = newY
    currentTetrominoSpin = newSpin
    return
  }
}

func isValidMove(x, y, spin int) bool {
  if x < 0 {
    return false
  }

  blocks := blocksOf(currentTetromino, x, y, spin)

  for _, block := range blocks {
    if block[1] >= 0 {
      if block[0] > (boardWidth - 1) || block[1] > (boardHeight - 1) {
        return false
      }

      if currentBoard[block[1]][block[0]] != termbox.ColorDefault {
        return false
      }
    }
  }

  return true
}

func isTouchingGround(x, y int) bool {
  blocks := blocksOf(currentTetromino, x, y, currentTetrominoSpin)

  for _, block := range blocks {
    if block[1] >= 0 {
      if block[1] + 1 >= boardHeight || currentBoard[block[1] + 1][block[0]] != termbox.ColorDefault {
        return true
      }
    }
  }

  return false
}

func blocksOf(t Tetromino, x, y, spin int) [4][2]int {
  normalizedSpin := spin % len(TetrominoShapes[currentTetromino])

  var blocks [4][2]int
  var i = 0
  for r, cols := range TetrominoShapes[currentTetromino][normalizedSpin] {
    for c, flag := range cols {
      if flag == 1 {
        blocks[i] = [2]int {x + c, y + r}
        i++
      }
    }
  }
  return blocks
}

func newRow() [boardWidth]termbox.Attribute {
  var row [boardWidth]termbox.Attribute
  for i := 0; i < boardWidth; i++ {
    row[i] = termbox.ColorDefault
  }

  return row
}

func clearFullRows() {
  var fullRows []int
  for r := 0; r < boardHeight; r++ {
    if isFullRow(r) {
      fullRows = append(fullRows, r)
    }
  }

  for _, r := range fullRows {
    currentBoard = append(currentBoard[:r], currentBoard[r+1:]...)
    currentBoard = append([][boardWidth]termbox.Attribute { newRow() }, currentBoard...)
  }
}

func isFullRow(r int) bool {
  for c := 0; c < boardWidth; c++ {
    if currentBoard[r][c] == termbox.ColorDefault {
      return false
    }
  }

  return true
}

func setScore(score string) {
  gameScore = score
}
