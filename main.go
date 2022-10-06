package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/riadafridishibly/go-snake/game"
)

func main() {
	ebiten.SetWindowSize(game.Width, game.Height)
	ebiten.SetWindowTitle("Snake! 🐍")
	ebiten.SetTPS(15)
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
