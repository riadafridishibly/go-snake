package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Dir uint8

const (
	None Dir = iota
	Up
	Down
	Left
	Right
)

type KeyboardInput struct{}

func (i *KeyboardInput) GetDirection() Dir {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		return Up
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		return Down
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		return Left
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		return Right
	}

	return None
}
