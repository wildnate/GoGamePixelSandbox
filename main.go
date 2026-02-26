package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	animation Animation
	img       *ebiten.Image
}

func (g *Game) Update() error {
	g.animation.Update()
	return nil

}

type Animation struct {
	First, Last, Cur, Speed, Elapsed int
}

func (a *Animation) Frame(SpriteSheetWidthTiles int) image.Rectangle {
	tx := a.Cur % SpriteSheetWidthTiles
	ty := a.Cur / SpriteSheetWidthTiles

	px := tx * tilesize
	py := ty * tilesize
	return image.Rect(px, py, px.tilesize, py.tilesize)
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(6, 6)
	opts.GeoM.Translate(50, 50)

	img := g.img.SubImage(g.animation.Frame(4)).(*ebiten.Image)
	screen.DrawImage(img, &opts)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (a *Animation) Update() {
	a.Elapsed += 1
	if a.Elapsed > a.Speed {
		a.Elapsed -= a.Speed
		a.Cur += 1
		if a.Cur > a.Last {
			a.Cur = a.First
		}
	}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	img, _, err := ebitenutil.NewImageFromFile("Spritesheet.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{img: img,
		animation: Animation{
			First: 0,
			Last:  4,
			Speed: 10,
		}}); err != nil {
		log.Fatal(err)
	}
}
