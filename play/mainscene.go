package play

import (
	"image/color"
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type SysList struct {
	RenderSys *common.RenderSystem
	DragSys   *DragSystem
	VelSys    *VelocitySystem
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

	for i := 0; i < 1000; i++ {
		a := NewBall(rand.Float32()*600, rand.Float32()*400, rand.Float32()*20)
		b := NewBall(rand.Float32()*600, rand.Float32()*400, rand.Float32()*20)
		sList.DragSys.Connect(a, b, 10)
		sList.RenderSys.AddByInterface(a)
		sList.RenderSys.AddByInterface(b)
		sList.VelSys.Add(a)
		sList.VelSys.Add(b)

	}

	w.AddSystem(sList.RenderSys)
	w.AddSystem(sList.DragSys)
	w.AddSystem(sList.VelSys)

}
