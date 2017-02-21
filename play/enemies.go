package play

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type Boxy struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	VelocityComponent
}

type BoxSystem struct {
	boxes   []Boxy
	targets []Spaceface
}
