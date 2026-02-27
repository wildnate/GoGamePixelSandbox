package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var PlayerXMain, PlayerYMain float64

type Sprite struct {
	Img      *ebiten.Image
	X, Y     float64
	tilesize float64
}
type Enemy struct {
	Img           *ebiten.Image
	X, Y          float64
	tilesize      float64
	FollowsPlayer bool
	anim          Animation
}
type Player struct {
	playerImage *ebiten.Image
	X, Y        float64
	tilesize    float64
	anim        Animation
}

type Game struct {
	animation Animation
	img       *ebiten.Image
	Player    *Player
	sprites   []*Sprite
	enemies   []*Enemy
}
type Animation struct {
	First, Last, Cur, Speed, Elapsed int
}

func (g *Game) Update() error {
	g.Player.anim.Update()
	for _, e := range g.enemies {
		e.anim.Update()
		if e.FollowsPlayer {
			// chase logic, etc.
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.X += 2
		PlayerXMain = g.Player.X
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.X -= 2
		PlayerXMain = g.Player.X
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.Y += 2
		PlayerYMain = g.Player.Y
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.Y -= 2
		PlayerYMain = g.Player.Y
	}

	return nil

}

// PlayerTileSize is the width/height of a single animation frame in pixels.
// Adjust this to match your spritesheet.

func (a *Animation) Frame(spriteSheetWidthTiles, tilesize int) image.Rectangle {
	tx := a.Cur % spriteSheetWidthTiles
	ty := a.Cur / spriteSheetWidthTiles

	px := tx * tilesize
	py := ty * tilesize
	// bottom-right coordinates are start + tilesize
	return image.Rect(px, py, px+tilesize, py+tilesize)
}

func (g *Game) Draw(screen *ebiten.Image) {

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.Player.Y, g.Player.X)

	screen.DrawImage(
		g.Player.playerImage.
			SubImage(g.Player.anim.Frame(64, 96)).(*ebiten.Image),
		&opts,
	)

	for _, e := range g.enemies {
		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(e.X, e.Y)
		screen.DrawImage(
			e.Img.SubImage(e.anim.Frame(64, 96)).(*ebiten.Image),
			&opts,
		)
	}
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
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Game Main")
	playerImage, _, err := ebitenutil.NewImageFromFile("assets/player/player-idle-96x96.png")
	skeletonImage, _, err := ebitenutil.NewImageFromFile("assets/enemies/Skeleton.png")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{

		enemies: []*Enemy{
			{

				Img:           skeletonImage,
				X:             100.0,
				Y:             100.0,
				tilesize:      0,
				FollowsPlayer: false,
				anim:          Animation{First: 0, Last: 7, Speed: 10},
			},
			{
				Img:           skeletonImage,
				X:             150.0,
				Y:             50.0,
				tilesize:      0,
				FollowsPlayer: true,
				anim:          Animation{First: 0, Last: 7, Speed: 10},
			},
		},
		Player: &Player{
			tilesize:    0,
			playerImage: playerImage,
			X:           100,
			Y:           100,
			anim:        Animation{First: 0, Last: 3, Speed: 6},
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}
