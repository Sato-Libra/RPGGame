package systems

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type PlayerSystem struct {
	world        *ecs.World
	playerEntity *Player
	texture      *common.Texture
}

func (ps *PlayerSystem) New(w *ecs.World) {
	ps.world = w
	// プレーヤーの作成
	player := Player{BasicEntity: ecs.NewBasic()}

	// 初期の配置
	positionX := int(engo.WindowWidth() / 2)
	positionY := int(engo.WindowHeight() - 88)
	player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: float32(positionX), Y: float32(positionY)},
		Width:    30,
		Height:   30,
	}
	// 画像の読み込み
	texture, err := common.LoadedSprite("pics/greenoctocat.png")
	if err != nil {
		fmt.Println("Unable to load texture: " + err.Error())
	}
	player.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: 0.1, Y: 0.1},
	}
	player.RenderComponent.SetZIndex(1)
	ps.playerEntity = &player
	ps.texture = texture
	for _, system := range ps.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)
		}
	}
	common.CameraBounds = engo.AABB{
		Min: engo.Point{X: 0, Y: 0},
		Max: engo.Point{X: 40000, Y: 300},
	}
}

func (*PlayerSystem) Remove(ecs.BasicEntity) {}

func (ps *PlayerSystem) Update(dt float32) {
}
