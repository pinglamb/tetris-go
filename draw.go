package main

import "fmt"
import "github.com/nsf/termbox-go"

const cellWidth = 5
const cellHeight = 2

const boardWidth = 10
const boardHeight = 16

const topBorderWidth = 1
const leftBorderWidth = 2
const topPadding = 1
const leftPadding = 2

const boardOffsetX = 10
const boardOffsetY = 5
const nextPaneOffsetX = boardOffsetX + boardWidth * cellWidth + 2 * leftBorderWidth + leftPadding * 2 + 5
const nextPaneOffsetY = boardOffsetY
const holdPaneOffsetX = nextPaneOffsetX
const holdPaneOffsetY = nextPaneOffsetY + 4 * cellHeight + 2 * topBorderWidth + 2 * topPadding + 3
const scorePaneOffsetX = nextPaneOffsetX
const scorePaneOffsetY = holdPaneOffsetY + 4 * cellHeight + 2 * topBorderWidth + 2 * topPadding + 3

const peerNextPaneOffsetX = 120
const peerNextPaneOffsetY = boardOffsetY
const peerHoldPaneOffsetX = peerNextPaneOffsetX
const peerHoldPaneOffsetY = peerNextPaneOffsetY + 4 * cellHeight + 2 * topBorderWidth + 2 * topPadding + 3
const peerScorePaneOffsetX = peerNextPaneOffsetX
const peerScorePaneOffsetY = peerHoldPaneOffsetY + 4 * cellHeight + 2 * topBorderWidth + 2 * topPadding + 3
const peerBoardOffsetX = peerNextPaneOffsetX + 4 * cellWidth + 2 * leftBorderWidth + leftPadding * 2 + 5
const peerBoardOffsetY = boardOffsetY

var log string

func drawGame() {
  drawInstructions()
  drawPanes()

  if gameStarted {
    drawBoard(boardOffsetX, boardOffsetY, currentBoard)
    if !dead {
      drawTetrominoOnBoard(boardOffsetX, boardOffsetY, currentTetromino, currentTetrominoSpin, currentTetrominoX, currentTetrominoY)
    }
    drawTetrominoOnNextPane(nextPaneOffsetX, nextPaneOffsetY, nextTetromino)
    if hasTetrominoHolded {
      drawTetrominoOnHoldPane(holdPaneOffsetX, holdPaneOffsetY, holdedTetromino)
    }
  }

  if isMP() {
    drawPeerGame()
  }

  drawScore(scorePaneOffsetX, scorePaneOffsetY, gameScore)
}

func setLog(t string) {
  log = t
}

func drawText(x, y int, text string) {
  for i, ch := range text {
    termbox.SetCell(x + i, y, ch, termbox.ColorWhite, termbox.ColorBlack)
  }
}

func drawInstructions() {
  _, h := termbox.Size()
  shift := 9
  drawText(boardOffsetX, h - shift, "n: start/restart the game")
  drawText(boardOffsetX, h - shift + 1, "a, d: left/right")
  drawText(boardOffsetX, h - shift + 2, "w: drop")
  drawText(boardOffsetX, h - shift + 3, "j: spin")
  drawText(boardOffsetX, h - shift + 4, "h: hold")
  drawText(boardOffsetX, h - shift + 5, fmt.Sprintf("[Your IP]:%s to connect", myPort))
  drawText(boardOffsetX, h - shift + 6, fmt.Sprintf("Peer: %s", peerInfo))
  drawText(boardOffsetX, h - shift + 7, fmt.Sprintf("Log: %s", log))
}

func drawPanes() {
  drawPane(boardOffsetX, boardOffsetY, boardWidth, boardHeight, " Player 1 ")
  drawPane(nextPaneOffsetX, nextPaneOffsetY, 4, 4, " Next ")
  drawPane(holdPaneOffsetX, holdPaneOffsetY, 4, 4, " Hold ")
  drawPane(scorePaneOffsetX, scorePaneOffsetY, 4, 1, " Score ")
}

func drawPane(x, y, width, height int, title string) {
  color := termbox.ColorWhite
  for i := 0; i < width * cellWidth + 2 * leftBorderWidth + leftPadding * 2; i++ {
    termbox.SetCell(x + i, y, ' ', color, color)
    termbox.SetCell(x + i, y + height * cellHeight + topBorderWidth + 2 * topPadding, ' ', color, color)
  }

  for i := 0; i < height * cellHeight + topBorderWidth + 2 * topPadding; i++ {
    for j := 0; j < leftBorderWidth; j++ {
      termbox.SetCell(x + j, y + i, ' ', color, color)
      termbox.SetCell(x + width * cellWidth + leftBorderWidth + leftPadding * 2 + j, y + i, ' ', color, color)
    }
  }

  for i, ch := range title {
    termbox.SetCell(x + i + 4, y, ch, termbox.ColorWhite, termbox.ColorBlack)
  }
}

func drawBoard(offsetX, offsetY int, board [][boardWidth]termbox.Attribute) {
  for i := 0; i < boardHeight; i++ {
    for j := 0; j < boardWidth; j++ {
      drawPoint(j, i, offsetX + leftBorderWidth + leftPadding, offsetY + topBorderWidth + topPadding, board[i][j])
    }
  }
}

func drawTetrominoOnBoard(offsetX, offsetY int, t Tetromino, spin int, x, y int) {
  drawTetromino(t, spin, x, y, offsetX + leftBorderWidth + leftPadding, offsetY + topBorderWidth + topPadding)
}

func drawTetrominoOnNextPane(offsetX, offsetY int, t Tetromino) {
  drawTetromino(t, 0, 1, 0, offsetX + leftBorderWidth + leftPadding, offsetY + topBorderWidth + topPadding)
}

func drawTetrominoOnHoldPane(offsetX, offsetY int, t Tetromino) {
  drawTetromino(t, 0, 0, 0, offsetX + leftBorderWidth + leftPadding, offsetY + topBorderWidth + topPadding)
}

func drawScore(offsetX, offsetY int, score string) {
  for i, ch := range score {
    termbox.SetCell(offsetX + leftBorderWidth + leftPadding + i + 4, offsetY + topBorderWidth + topPadding, ch, termbox.ColorWhite, termbox.ColorBlack)
  }
}

func drawTetromino(t Tetromino, spin int, x, y, offsetX, offsetY int) {
  color := TetrominoColors[t]
  normalizedSpin := spin % len(TetrominoShapes[t])
  for r, cols := range TetrominoShapes[t][normalizedSpin] {
    for c, flag := range cols {
      if flag == 1 {
        drawPoint(x + c, y + r, offsetX, offsetY, color)
      }
    }
  }
}

func drawPoint(x, y, offsetX, offsetY int, color termbox.Attribute) {
  for i := 0; i < cellWidth; i++ {
    for j := 0; j < cellHeight; j++ {
      termbox.SetCell(offsetX + x * cellWidth + i, offsetY + y * cellHeight + j, ' ', color, color)
    }
  }
}

func drawPeerGame() {
  drawPane(peerBoardOffsetX, peerBoardOffsetY, boardWidth, boardHeight, " Player 2 ")
  drawPane(peerNextPaneOffsetX, peerNextPaneOffsetY, 4, 4, " Next ")
  drawPane(peerHoldPaneOffsetX, peerHoldPaneOffsetY, 4, 4, " Hold ")
  drawPane(peerScorePaneOffsetX, peerScorePaneOffsetY, 4, 1, " Score ")

  if gameStarted {
    drawBoard(peerBoardOffsetX, peerBoardOffsetY, peerCurrentBoard)
    if !peerDead {
      drawTetrominoOnBoard(peerBoardOffsetX, peerBoardOffsetY, peerCurrentTetromino, peerCurrentTetrominoSpin, peerCurrentTetrominoX, peerCurrentTetrominoY)
    }
    drawTetrominoOnNextPane(peerNextPaneOffsetX, peerNextPaneOffsetY, peerNextTetromino)
    if peerHasTetrominoHolded {
      drawTetrominoOnHoldPane(peerHoldPaneOffsetX, peerHoldPaneOffsetY, peerHoldedTetromino)
    }
  }
}
