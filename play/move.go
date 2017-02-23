package play

import (
	"image/color"

	"github.com/coderconvoy/engotil"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type BallSystem struct {
}

type BoySystem struct {
}

func PlayerColor(n int) color.Color {
	switch n {
	case 0:
		return color.RGBA{50, 0, 0, 255}
	default:
		return color.RGBA{0, 0, 50, 255}
	}
}

type Boy struct {
	ecs.BasicEntity
	DragComponent
	engotil.VelocityComponent
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
	ControlComponent
}

func NewBoy(x, y, w float32, pnum int) *Boy {
	res := &Boy{
		BasicEntity:       ecs.NewBasic(),
		DragComponent:     DragComponent{w},
		VelocityComponent: engotil.VelocityComponent{Friction: 10},
		RenderComponent: common.RenderComponent{
			Drawable: common.Triangle{},
			Color:    PlayerColor(pnum),
		},
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{x, y},
			Width:    w,
			Height:   w,
		},
		ControlComponent: GetKeys(pnum),
		CollisionComponent: common.CollisionComponent{
			Solid: true,
			Main:  true,
		},
	}
	res.SetZIndex(5)
	return res

}

type Ball struct {
	ecs.BasicEntity
	DragComponent
	engotil.VelocityComponent
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
}

func NewBall(x, y, w float32, pnum int) *Ball {
	res := &Ball{
		BasicEntity:       ecs.NewBasic(),
		DragComponent:     DragComponent{w},
		VelocityComponent: engotil.VelocityComponent{Friction: 0.5},
		RenderComponent: common.RenderComponent{
			Drawable: common.Circle{},
			Color:    PlayerColor(pnum),
		},
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{x, y},
			Width:    w,
			Height:   w,
		},
		CollisionComponent: common.CollisionComponent{
			Main:  true,
			Solid: false,
		},
	}
	res.SetZIndex(5)
	return res
}
