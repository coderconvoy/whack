package play

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/coderconvoy/engotil"
)

type DragComponent struct {
	weight float32
}

func (dc *DragComponent) GetDragComponent() *DragComponent {
	return dc
}

type Draggable interface {
	GetBasicEntity() *ecs.BasicEntity
	GetSpaceComponent() *common.SpaceComponent
	Push(engo.Point)
	GetDragComponent() *DragComponent
}

type Teather struct {
	ecs.BasicEntity
	a, b    Draggable
	tension float32
	length  float32
}

type DragSystem struct {
	teathers []*Teather
}

func (ds *DragSystem) Connect(a, b Draggable, tension, length float32) *Teather {
	res := &Teather{
		BasicEntity: ecs.NewBasic(),
		a:           a,
		b:           b,
		tension:     tension,
		length:      length,
	}

	ds.teathers = append(ds.teathers, res)
	return res
}

func (ds *DragSystem) Remove(b ecs.BasicEntity) {
}

func (ds *DragSystem) Update(d float32) {
	for _, t := range ds.teathers {
		aloc := t.a.GetSpaceComponent()
		acen := engotil.SpaceCenter(*aloc)

		bloc := t.b.GetSpaceComponent()
		bcen := engotil.SpaceCenter(*bloc)

		adrag := t.a.GetDragComponent()
		bdrag := t.b.GetDragComponent()
		d2 := aloc.Position.PointDistance(bloc.Position)
		if d2 < t.length {
			continue
		}

		tension := t.tension * (d2 - t.length)
		combi := adrag.weight + bdrag.weight
		if combi == 0 {
			combi = 0.01
		}
		aFrac := tension * d * bdrag.weight / combi
		t.a.Push(engo.Point{
			aFrac * (bcen.X - acen.X),
			aFrac * (bcen.Y - acen.Y),
		})
		bFrac := tension * d * adrag.weight / combi
		t.b.Push(engo.Point{
			bFrac * (acen.X - bcen.X),
			bFrac * (acen.Y - bcen.Y),
		})
	}
}
