package play

import (
	"engo.io/ecs"
	"engo.io/engo"
)

type ControlComponent struct {
	ctrls []KeyCommand
}

func (cc *ControlComponent) GetControls() []KeyCommand {
	return cc.ctrls
}

type KeyCommand struct {
	K    engo.Key
	S    string
	Push engo.Point
}

func GetKeys(n int) ControlComponent {
	if n == 0 {
		return ControlComponent{[]KeyCommand{
			{engo.ArrowLeft, "l1", engo.Point{-1, 0}},
			{engo.ArrowRight, "r1", engo.Point{1, 0}},
			{engo.ArrowUp, "u1", engo.Point{0, -1}},
			{engo.ArrowDown, "d1", engo.Point{0, 1}},
		}}
	}
	return ControlComponent{[]KeyCommand{
		{engo.A, "l2", engo.Point{-1, 0}},
		{engo.D, "r2", engo.Point{1, 0}},
		{engo.W, "u2", engo.Point{0, -1}},
		{engo.S, "d2", engo.Point{0, 1}},
	}}
}

type Controlable interface {
	ID() uint64
	GetControls() []KeyCommand
	Push(engo.Point)
}

type ControlSystem struct {
	obs []Controlable
}

func (cs *ControlSystem) Add(c Controlable) {
	cs.obs = append(cs.obs, c)
}

func (cs *ControlSystem) Remove(ecs.BasicEntity) {
}

func (cs *ControlSystem) Update(d float32) {
	for _, o := range cs.obs {
		for _, c := range o.GetControls() {
			if engo.Input.Button(c.S).Down() {
				o.Push(c.Push)
			}
		}
	}
}
