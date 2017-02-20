package play

import (
	"image/color"

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
	VelocityComponent
	common.RenderComponent
	common.SpaceComponent
	ControlComponent
}

func NewBoy(x, y, w float32, pnum int) *Boy {
	return &Boy{
		BasicEntity:       ecs.NewBasic(),
		DragComponent:     DragComponent{w},
		VelocityComponent: VelocityComponent{friction: 10},
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
	}

}

type Ball struct {
	ecs.BasicEntity
	DragComponent
	VelocityComponent
	common.RenderComponent
	common.SpaceComponent
}

func NewBall(x, y, w float32) *Ball {
	return &Ball{
		BasicEntity:       ecs.NewBasic(),
		DragComponent:     DragComponent{w},
		VelocityComponent: VelocityComponent{friction: 0.1},
		RenderComponent: common.RenderComponent{
			Drawable: common.Circle{},
			Color:    color.Black,
		},
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{x, y},
			Width:    w,
			Height:   w,
		},
	}
}
