package main

import (
	"bytes"
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/KMimura/RPGGame/systems"
	"golang.org/x/image/font/gofont/gosmallcaps"
)
//gogo
type myScene struct{}

func (*myScene) Type() string { return "myGame" }

func (*myScene) Preload() {
	engo.Files.Load("pics/greenoctocat_top.png", "pics/greenoctocat_left.png", "pics/greenoctocat_right.png", "pics/greenoctocat_bottom.png", "pics/ghost.png", "pics/overworld_tileset_grass.png", "pics/stone_dot.png")
	engo.Files.LoadReaderData("go.ttf", bytes.NewReader(gosmallcaps.TTF))
	common.SetBackground(color.RGBA{255, 250, 220, 0})
}

func (*myScene) Setup(u engo.Updater) {
	engo.Input.RegisterButton("MoveRight", engo.KeyD, engo.KeyArrowRight)
	engo.Input.RegisterButton("MoveLeft", engo.KeyA, engo.KeyArrowLeft)
	engo.Input.RegisterButton("MoveUp", engo.KeyW, engo.KeyArrowUp)
	engo.Input.RegisterButton("MoveDown", engo.KeyS, engo.KeyArrowDown)
	engo.Input.RegisterButton("Space", engo.KeySpace)
	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.TileSystem{})
	world.AddSystem(&systems.PlayerSystem{})
	world.AddSystem(&systems.EnemySystem{})
}

func (*myScene) Exit() {
	engo.Exit()
}

func main() {
	opts := engo.RunOptions{
		Title:          "myGame",
		Width:          400,
		Height:         300,
		StandardInputs: true,
		NotResizable:   true,
	}
	engo.Run(opts, &myScene{})
}
