package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/riadafridishibly/go-snake/utils"
)

type Game struct {
	snake *Snake
	food  *Food
	dir   Dir
	input KeyboardInput
}

func (g *Game) Update() error {
	if dir := g.input.GetDirection(); dir != None {
		g.dir = dir
	}
	if g.dir != None {
		g.snake.Update(g.dir, g.food)
	}
	return nil
}

func (g *Game) drawBackground(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.Fill(color.RGBA{
		A: 255,
		R: 21,
		G: 21,
		B: 20,
	})
	for i := 0; i < Cols; i++ {
		for j := 0; j < Rows; j++ {
			utils.DrawCell(screen, i, j, color.Black, Gap, CellSize)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBackground(screen)
	g.snake.Draw(screen)
	g.food.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func NewGame() *Game {
	return &Game{
		snake: NewSnake(),
		food:  NewFood(),
	}
}
