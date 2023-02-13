package main

import (
	"ledsuit/internal/led"
	"ledsuit/internal/pattern"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	current_pattern int
	suit            *Suit
	pattern         [8]pattern.Pattern
	tick            uint32
}

func NewGame() *Game {
	suit := NewSuit("GR")

	var patterns = [8]pattern.Pattern{
		pattern.NewDarkPattern(suit),
		pattern.NewHellsBellsPattern(suit),
		pattern.NewAmmelieAltPattern(suit, led.ORANGE),
		pattern.NewKitPattern(suit),
		pattern.NewHueShiftPattern(suit),
		pattern.NewTetrisPattern(suit, 0),
		pattern.NewHeartPlusPattern(suit, led.WHITE),
		pattern.NewHeartPattern(suit),
	}

	return &Game{
		current_pattern: 0,
		suit:            suit,
		pattern:         patterns,
		tick:            0,
	}
}

func (g *Game) checkInput() {
	new_pattern := g.current_pattern
	if inpututil.IsKeyJustPressed(ebiten.Key0) {
		new_pattern = 0
	} else if inpututil.IsKeyJustPressed(ebiten.Key1) {
		new_pattern = 1
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		new_pattern = 2
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
		new_pattern = 3
	} else if inpututil.IsKeyJustPressed(ebiten.Key4) {
		new_pattern = 4
	} else if inpututil.IsKeyJustPressed(ebiten.Key5) {
		new_pattern = 5
	} else if inpututil.IsKeyJustPressed(ebiten.Key6) {
		new_pattern = 6
	} else if inpututil.IsKeyJustPressed(ebiten.Key7) {
		new_pattern = 7
	}

	if new_pattern != g.current_pattern {
		g.pattern[g.current_pattern].Reset()
		g.current_pattern = new_pattern
	}
}

func (g *Game) Update() error {
	g.checkInput()

	g.pattern[g.current_pattern].SetLEDs(g.tick)
	g.suit.Display()

	g.tick++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.suit.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 800
}

func main() {

	ebiten.SetWindowSize(800, 800)
	ebiten.SetTPS(17)
	ebiten.SetWindowTitle("Led suit test ground)")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
