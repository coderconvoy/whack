package play

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type VelocityComponent struct {
	engo.Point
	friction float32
}

func (vc *VelocityComponent) Push(f engo.Point) {
	vc.X += f.X
	vc.Y += f.Y
}

func (vc *VelocityComponent) GetVelocityComponent() *VelocityComponent {
	return vc
}

type Velocitable interface {
	ID() uint64
	GetVelocityComponent() *VelocityComponent
	GetSpaceComponent() *common.SpaceComponent
}

type VelocitySystem struct {
	obs []Velocitable
}

func (vs *VelocitySystem) Add(ob Velocitable) {
	vs.obs = append(vs.obs, ob)
}

func (vs *VelocitySystem) Update(d float32) {
	for _, o := range vs.obs {
		vc := o.GetVelocityComponent()
		vc.Point.MultiplyScalar(1 - (vc.friction * d))

		sc := o.GetSpaceComponent()
		sc.Position.X += vc.X
		sc.Position.Y += vc.Y

	}
}

func (vs *VelocitySystem) Remove(ob ecs.BasicEntity) {
}
