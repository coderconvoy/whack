package play

import (
	"image/color"
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type SysList struct {
	RenderSys  *common.RenderSystem
	DragSys    *DragSystem
	VelSys     *VelocitySystem
	ControlSys *ControlSystem
}

type MainScene struct{ NPlayers int }

func (*MainScene) Type() string { return "MainScene" }

func (*MainScene) Preload() {
}

func (ms *MainScene) Setup(w *ecs.World) {
	common.SetBackground(color.White)

	var sList SysList

	sList.RenderSys = &common.RenderSystem{}
	sList.DragSys = &DragSystem{}
	sList.VelSys = &VelocitySystem{}
	sList.ControlSys = &ControlSystem{}

	for i := 0; i < 3; i++ {

		a := NewBoy(rand.Float32()*600, rand.Float32()*400, 10+rand.Float32()*20, i)
		b := NewBall(rand.Float32()*600, rand.Float32()*400, 2+rand.Float32()*10)
		if i < 2 {
			for _, kc := range a.GetControls() {
				engo.Input.RegisterButton(kc.S, kc.K)
			}
			sList.ControlSys.Add(a)
		}
		sList.DragSys.Connect(a, b, 1, 30)
		sList.RenderSys.AddByInterface(a)
		sList.RenderSys.AddByInterface(b)
		sList.VelSys.Add(a)
		sList.VelSys.Add(b)

	}

	w.AddSystem(sList.RenderSys)
	w.AddSystem(sList.ControlSys)
	w.AddSystem(sList.DragSys)
	w.AddSystem(sList.VelSys)

}
