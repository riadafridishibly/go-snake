package game

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/riadafridishibly/go-snake/utils"
)

var snakeColor = color.RGBA{
	R: 180,
	G: 100,
	B: 58,
	A: 255,
}

var DirX = [...]int{
	None:  0,
	Up:    0,
	Down:  0,
	Left:  -1,
	Right: +1,
}

var DirY = [...]int{
	None:  0,
	Up:    -1,
	Down:  +1,
	Left:  0,
	Right: 0,
}

type Cell struct {
	x int
	y int
	c color.Color
}

func (c Cell) Draw(screen *ebiten.Image) {
	utils.DrawCell(screen, c.x, c.y, c.c, Gap, CellSize)
}

type Snake struct {
	cells []Cell
}

func (s *Snake) foodInsideSnake(x, y int) bool {
	for i := range s.cells {
		if s.cells[i].x == x && s.cells[i].y == y {
			return true
		}
	}
	return false
}

func (s *Snake) updateFoodPosition(food *Food) {
	x := 0
	y := 0
	for i := 0; i < 5; i++ {
		x = rand.Intn(Cols)
		y = rand.Intn(Rows)
		if !s.foodInsideSnake(x, y) {
			break
		}
	}

	food.Update(x, y)
}

func (s *Snake) CheckCollision(food *Food, dir Dir) bool {
	if food.x == s.cells[0].x && food.y == s.cells[0].y {
		s.cells = append(s.cells, Cell{})
		s.update(dir)
		s.updateFoodPosition(food)
	}
	return false
}

func (s *Snake) update(dir Dir) {
	// Copy all the blocks
	for i := len(s.cells) - 1; i > 0; i-- {
		s.cells[i].x = s.cells[i-1].x
		s.cells[i].y = s.cells[i-1].y
		s.cells[i].c = s.cells[i-1].c
	}

	s.cells[0].x += DirX[dir]
	s.cells[0].y += DirY[dir]

	if s.cells[0].x < 0 {
		s.cells[0].x = Cols - 1
	}
	if s.cells[0].x >= Cols {
		s.cells[0].x = 0
	}
	if s.cells[0].y < 0 {
		s.cells[0].y = Rows - 1
	}
	if s.cells[0].y >= Rows {
		s.cells[0].y = 0
	}
}

func (s *Snake) Update(dir Dir, food *Food) {
	s.update(dir)
	s.CheckCollision(food, dir)
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for i := range s.cells {
		s.cells[i].Draw(screen)
	}
}

func NewSnake() *Snake {
	return &Snake{
		cells: []Cell{
			{
				x: 4,
				y: 3,
				c: snakeColor,
			},
			{
				x: 3,
				y: 3,
				c: snakeColor,
			},
			{
				x: 2,
				y: 3,
				c: snakeColor,
			},
		},
	}
}
