package utils

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawCell(screen *ebiten.Image, x, y int, c color.Color, gap, cellSize int) {
	xPos := float64(x*cellSize + gap)
	yPos := float64(y*cellSize + gap)
	ebitenutil.DrawRect(
		screen,
		xPos, yPos,
		float64(cellSize-2*gap), float64(cellSize-2*gap), c)
}
