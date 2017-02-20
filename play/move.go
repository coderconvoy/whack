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
	common.RenderComponent
	common.SpaceComponent
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
		VelocityComponent: VelocityComponent{friction: 0.2},
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
