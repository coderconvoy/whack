package play

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/coderconvoy/engotil"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type SysList struct {
	RenderSys    *common.RenderSystem
	DragSys      *DragSystem
	VelSys       *engotil.VelocitySystem
	ControlSys   *ControlSystem
	BoxSys       *BoxSystem
	CollisionSys *engotil.CollisionSystem
}

type MainScene struct{ NPlayers int }

func (*MainScene) Type() string { return "MainScene" }

func (*MainScene) Preload() {
	err := engo.Files.Load("lev1.tmx")
	if err != nil {
		fmt.Println("No Load lev1.tmx : ", err)
	}
}

func (ms *MainScene) Setup(w *ecs.World) {
	common.SetBackground(color.White)

	var sList SysList

	sList.RenderSys = &common.RenderSystem{}
	sList.DragSys = &DragSystem{}
	sList.VelSys = &engotil.VelocitySystem{}
	sList.ControlSys = &ControlSystem{}
	sList.BoxSys = &BoxSystem{}
	sList.CollisionSys = &engotil.CollisionSystem{}

	_ = LoadMap("lev1.tmx", sList)

	for i := 0; i < ms.NPlayers; i++ {
		sx := 100 + rand.Float32()*400
		sy := 60 + rand.Float32()*280

		a := NewBoy(sx, sy, 10+rand.Float32()*20, i)
		b := NewBall(sx+rand.Float32()*50, sy+rand.Float32()*50, 5+rand.Float32()*10)
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
		sList.CollisionSys.Add(b)
		sList.CollisionSys.Add(a)

	}

	w.AddSystem(sList.RenderSys)
	w.AddSystem(sList.ControlSys)
	w.AddSystem(sList.DragSys)
	w.AddSystem(sList.VelSys)
	w.AddSystem(sList.BoxSys)
	w.AddSystem(sList.CollisionSys)
	w.AddSystem(&HitSystem{NPlayers: ms.NPlayers})
	w.AddSystem(NewSpawnSystem(sList))

}
