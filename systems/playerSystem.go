package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Player プレーヤーを表す構造体
type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	// 向き (0 => 上, 1 => 右, 2 => 下, 3 => 左)
	direction int
}

// PlayerSystem プレーヤーシステム
type PlayerSystem struct {
	world        *ecs.World
	playerEntity *Player
	texture      *common.Texture
}

var playerInstance *Player

// それぞれの向きのプレーヤーの画像
var topPic *common.Texture
var rightPic *common.Texture
var bottomPic *common.Texture
var leftPic *common.Texture

func (ps *PlayerSystem) New(w *ecs.World) {
	ps.world = w
	// プレーヤーの作成
	player := Player{BasicEntity: ecs.NewBasic()}

	playerInstance = &player

	// 初期の配置
	positionX := int(engo.WindowWidth() / 2)
	positionY := int(engo.WindowHeight() - 88)
	player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: float32(positionX), Y: float32(positionY)},
		Width:    30,
		Height:   30,
	}
	// 画像の読み込み
	topPic, _ = common.LoadedSprite("pics/greenoctocat_top.png")
	rightPic, _ = common.LoadedSprite("pics/greenoctocat_right.png")
	bottomPic, _ = common.LoadedSprite("pics/greenoctocat_bottom.png")
	leftPic, _ = common.LoadedSprite("pics/greenoctocat_left.png")

	player.RenderComponent = common.RenderComponent{
		Drawable: topPic,
		Scale:    engo.Point{X: 0.1, Y: 0.1},
	}
	player.RenderComponent.SetZIndex(1)
	ps.playerEntity = &player
	ps.texture = topPic
	for _, system := range ps.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)
		}
	}
	common.CameraBounds = engo.AABB{
		Min: engo.Point{X: 0, Y: 0},
		Max: engo.Point{X: 4000, Y: 4000},
	}
}

// Remove 削除する
func (*PlayerSystem) Remove(ecs.BasicEntity) {}

// Update アップデートする
func (ps *PlayerSystem) Update(dt float32) {
	camX := camEntity.X()
	camY := camEntity.Y()
	if engo.Input.Button("MoveRight").Down() {
		if ps.playerEntity.direction != 1 {
			ps.playerEntity.direction = 1
		} else {
			ps.playerEntity.direction = 1
			if camX < 4000 {
				ps.playerEntity.SpaceComponent.Position.X += 5
				if ps.playerEntity.SpaceComponent.Position.X-camX > 100 {
					engo.Mailbox.Dispatch(common.CameraMessage{
						Axis:        common.XAxis,
						Value:       5,
						Incremental: true,
					})
				}
			}
		}
	} else if engo.Input.Button("MoveLeft").Down() {
		if ps.playerEntity.direction != 3 {
			ps.playerEntity.direction = 3
		} else {
			ps.playerEntity.direction = 3
			if camX > 200 {
				ps.playerEntity.SpaceComponent.Position.X -= 5
				if camX-ps.playerEntity.SpaceComponent.Position.X > 100 {
					engo.Mailbox.Dispatch(common.CameraMessage{
						Axis:        common.XAxis,
						Value:       -5,
						Incremental: true,
					})
				}
			} else if ps.playerEntity.SpaceComponent.Position.X > 5 {
				ps.playerEntity.SpaceComponent.Position.X -= 5
			}
		}
	} else if engo.Input.Button("MoveUp").Down() {
		if ps.playerEntity.direction != 0 {
			ps.playerEntity.direction = 0
		} else {
			ps.playerEntity.direction = 0
			if camY > 200 {
				ps.playerEntity.SpaceComponent.Position.Y -= 5
				if camY-ps.playerEntity.SpaceComponent.Position.Y > 100 {
					engo.Mailbox.Dispatch(common.CameraMessage{
						Axis:        common.YAxis,
						Value:       -5,
						Incremental: true,
					})
				}
			} else if ps.playerEntity.SpaceComponent.Position.Y > 5 {
				ps.playerEntity.SpaceComponent.Position.Y -= 5
			}
		}
	} else if engo.Input.Button("MoveDown").Down() {
		if ps.playerEntity.direction != 2 {
			ps.playerEntity.direction = 2
		} else {
			ps.playerEntity.direction = 2
			if camY < 4000 {
				ps.playerEntity.SpaceComponent.Position.Y += 5
				if ps.playerEntity.SpaceComponent.Position.Y-camY > 100 {
					engo.Mailbox.Dispatch(common.CameraMessage{
						Axis:        common.YAxis,
						Value:       5,
						Incremental: true,
					})
				}
			}
		}
	} else if engo.Input.Button("Space").JustPressed() {
		ps.world.AddSystem(&BulletSystem{})
	}
	switch ps.playerEntity.direction {
	case 0:
		ps.playerEntity.RenderComponent.Drawable = topPic
	case 1:
		ps.playerEntity.RenderComponent.Drawable = rightPic
	case 2:
		ps.playerEntity.RenderComponent.Drawable = bottomPic
	case 3:
		ps.playerEntity.RenderComponent.Drawable = leftPic
	}
}
