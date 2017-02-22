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
	return &Boy{
		BasicEntity:       ecs.NewBasic(),
		DragComponent:     DragComponent{w},
		VelocityComponent: engotil.VelocityComponent{Friction: 10},
		RenderComponent: common.RenderComponent{
			Drawable: common.Triangle{},
			Color:    color.Black,
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

}

type Ball struct {
	ecs.BasicEntity
	DragComponent
	engotil.VelocityComponent
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
}

func NewBall(x, y, w float32) *Ball {
	return &Ball{
		BasicEntity:       ecs.NewBasic(),
		DragComponent:     DragComponent{w},
		VelocityComponent: engotil.VelocityComponent{Friction: 0.5},
		RenderComponent: common.RenderComponent{
			Drawable: common.Circle{},
			Color:    color.Black,
		},
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{x, y},
			Width:    w,
			Height:   w,
		},
		CollisionComponent: common.CollisionComponent{
			Main:  true,
			Solid: true,
		},
	}
}
