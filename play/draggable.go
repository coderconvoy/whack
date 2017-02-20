package play

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
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
}

type DragSystem struct {
	teathers []*Teather
}

func (ds *DragSystem) Connect(a, b Draggable, tension float32) *Teather {
	res := &Teather{
		BasicEntity: ecs.NewBasic(),
		a:           a,
		b:           b,
		tension:     tension,
	}

	ds.teathers = append(ds.teathers, res)
	return res
}

func (ds *DragSystem) Remove(b ecs.BasicEntity) {
}

func (ds *DragSystem) Update(d float32) {
	for _, t := range ds.teathers {
		aloc := t.a.GetSpaceComponent()
		bloc := t.b.GetSpaceComponent()
		adrag := t.a.GetDragComponent()
		bdrag := t.b.GetDragComponent()
		combi := adrag.weight + bdrag.weight
		if combi == 0 {
			combi = 0.01
		}
		aFrac := t.tension * d * adrag.weight / combi
		t.a.Push(engo.Point{
			aFrac * (bloc.Position.X - aloc.Position.X),
			aFrac * (bloc.Position.Y - aloc.Position.Y),
		})
		bFrac := d * bdrag.weight / combi
		t.b.Push(engo.Point{
			bFrac * (aloc.Position.X - bloc.Position.X),
			bFrac * (aloc.Position.Y - bloc.Position.Y),
		})
	}
}
