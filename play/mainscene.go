package play

import (
	"image/color"
	"math/rand"

	"github.com/coderconvoy/engotil"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type SysList struct {
	RenderSys  *common.RenderSystem
	DragSys    *DragSystem
	VelSys     *engotil.VelocitySystem
	ControlSys *ControlSystem
	BoxSys     *BoxSystem
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
	sList.VelSys = &engotil.VelocitySystem{}
	sList.ControlSys = &ControlSystem{}
	sList.BoxSys = &BoxSystem{}

	for i := 0; i < 2; i++ {

		a := NewBoy(rand.Float32()*600, rand.Float32()*400, 10+rand.Float32()*20, i)
		b := NewBall(rand.Float32()*600, rand.Float32()*400, 2+rand.Float32()*10)
		if i < 2 {
			for _, kc := range a.GetControls() {
				engo.Input.RegisterButton(kc.S, kc.K)
			}
			sList.ControlSys.Add(a)
		}
		sList.DragSys.Connect(a, b, 0.1, 70)
		sList.RenderSys.AddByInterface(a)
		sList.RenderSys.AddByInterface(b)
		sList.VelSys.Add(a)
		sList.VelSys.Add(b)
		sList.BoxSys.AddTarget(a)

	}

	for i := 0; i < 10; i++ {
		b := NewBoxy(rand.Float32()*600, rand.Float32()*400)
		sList.RenderSys.AddByInterface(b)
		sList.VelSys.Add(b)
		sList.BoxSys.AddBox(b)
	}

	w.AddSystem(sList.RenderSys)
	w.AddSystem(sList.ControlSys)
	w.AddSystem(sList.DragSys)
	w.AddSystem(sList.VelSys)
	w.AddSystem(sList.BoxSys)

}
