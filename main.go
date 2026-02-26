package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}
type Enemy struct {
	*Sprite
	FollowsPlayer bool
}

type Game struct {
	animation        Animation
	img              *ebiten.Image
	playerX, PlayerY float64
	playerImage      *ebiten.Image
	player           *Sprite
	sprites          []*Sprite
	enemies          []*Enemy
}

func (g *Game) Update() error {
	g.animation.Update()
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.PlayerY += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.PlayerY -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerX += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerX -= 2
	}

	return nil

}

type Animation struct {
	First, Last, Cur, Speed, Elapsed int
}

// PlayerTileSize is the width/height of a single animation frame in pixels.
// Adjust this to match your spritesheet.
const PlayerTileSize = 96

func (a *Animation) Frame(spriteSheetWidthTiles int) image.Rectangle {
	tx := a.Cur % spriteSheetWidthTiles
	ty := a.Cur / spriteSheetWidthTiles

	px := tx * PlayerTileSize
	py := ty * PlayerTileSize
	// bottom-right coordinates are start + tilesize
	return image.Rect(px, py, px+PlayerTileSize, py+PlayerTileSize)
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	opts := ebiten.DrawImageOptions{}
	//opts.GeoM.Scale(6, 6)
	opts.GeoM.Translate(g.playerX, g.PlayerY)
	screen.DrawImage(
		g.playerImage.SubImage(g.animation.Frame(4)).(*ebiten.Image),
		&opts,
	)

	opts.GeoM.Reset()
	for _, sprite := range g.enemies {
		opts.GeoM.Translate(sprite.X, sprite.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 128, 128),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
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
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, To My Game")
	playerImage, _, err := ebitenutil.NewImageFromFile("assets/player/player-idle-96x96.png")
	skeletonImage, _, err := ebitenutil.NewImageFromFile("assets/Skeleton_Sword/Skeleton_White/Skeleton_With_VFX/SkeletonIdle.png")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		enemies: []*Enemy{
			{
				&Sprite{
					Img: skeletonImage,
					X:   100.0,
					Y:   100.0,
				},
				true,
			},
			{
				&Sprite{
					Img: skeletonImage,
					X:   150.0,
					Y:   50.0,
				},
				false,
			},
		},
		animation: Animation{
			First: 0, // first frame index
			Last:  3, // last frame index (0‑based)
			Speed: 8, // advance every 8 updates; tweak to taste
		},
		playerImage: playerImage,

		player: &Sprite{
			Img: playerImage,
			X:   100,
			Y:   100,
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}
