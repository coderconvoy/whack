package play

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/coderconvoy/engotil"
)

type Boxy struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	engotil.VelocityComponent
}

func NewBoxy(x, y float32) *Boxy {
	return &Boxy{
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{x, y},
			Width:    30,
			Height:   30,
		},
		RenderComponent: common.RenderComponent{
			Drawable: common.Rectangle{},
			Color:    color.Black,
		},
		VelocityComponent: engotil.VelocityComponent{
			Friction: 2,
		},
	}
}

type BoxSystem struct {
	boxes   []*Boxy
	targets []engotil.Spaceable
}

func (bs *BoxSystem) AddTarget(t engotil.Spaceable) {
	bs.targets = append(bs.targets, t)
}

func (bs *BoxSystem) AddBox(b *Boxy) {
	bs.boxes = append(bs.boxes, b)
}
func (bs *BoxSystem) Update(d float32) {
	if len(bs.targets) == 0 {
		return
	}

	for _, b := range bs.boxes {
		//loop for nearest target
		bcen := engotil.SpaceCenter(*b.GetSpaceComponent())
		nearest := bs.targets[0]
		ncen := engotil.SpaceCenter(*nearest.GetSpaceComponent())
		nd2 := bcen.PointDistanceSquared(ncen)

		for _, t := range bs.targets {
			ncen = engotil.SpaceCenter(*t.GetSpaceComponent())
			td2 := bcen.PointDistanceSquared(ncen)
			if td2 < nd2 {
				nearest = t
				nd2 = td2
			}
		}

		//push towards nearest

		ncen = engotil.SpaceCenter(*nearest.GetSpaceComponent())
		if bcen.X > ncen.X {
			b.Push(engo.Point{-0.4, 0})
		}
		if bcen.X < ncen.X {
			b.Push(engo.Point{0.4, 0})
		}
		if bcen.Y > ncen.Y {
			b.Push(engo.Point{0, -0.4})
		}
		if bcen.Y < ncen.Y {
			b.Push(engo.Point{0, 0.4})
		}
	}
}
func (bs *BoxSystem) Remove(e ecs.BasicEntity) {
}
