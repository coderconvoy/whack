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
	CollisionSys *engotil.GCollisionSystem
	LookSys      *LookSystem
}

type MainScene struct{ NPlayers int }

func (*MainScene) Type() string { return "MainScene" }

func (*MainScene) Preload() {
	err := engo.Files.Load("lev1.tmx")
	if err != nil {
		fmt.Println("No Load lev1.tmx : ", err)
	}
	err = engo.Files.Load("rat.png", "rat2.png", "boy1.png", "boy2.png")
	if err != nil {
		fmt.Println("could not load images:", err)

	}
}

func (ms *MainScene) Setup(w *ecs.World) {
	common.SetBackground(color.RGBA{100, 255, 150, 255})

	var sList SysList

	sList.RenderSys = &common.RenderSystem{}
	sList.DragSys = &DragSystem{}
	sList.VelSys = &engotil.VelocitySystem{}
	sList.ControlSys = &ControlSystem{}
	sList.BoxSys = &BoxSystem{}
	sList.CollisionSys = &engotil.GCollisionSystem{Solids: C_BOY | C_BALL | C_ENEMY}
	sList.LookSys = &LookSystem{}

	_ = LoadMap("lev1.tmx", sList)

	for i := 0; i < ms.NPlayers; i++ {
		sx := 100 + rand.Float32()*500
		sy := 100 + rand.Float32()*500

		a := NewBoy(sx, sy, 20, i)
		if i < 2 {
			for _, kc := range a.GetControls() {
				engo.Input.RegisterButton(kc.S, kc.K)
			}
			sList.ControlSys.Add(a)
		}
		sList.RenderSys.AddByInterface(a)
		sList.VelSys.Add(a)
		sList.BoxSys.AddTarget(a)
		sList.CollisionSys.Add(a)
		b := AddBall(a, i, 0.1, 70, sList)
		sList.LookSys.Connect(a, b)

	}

	w.AddSystem(sList.RenderSys)
	w.AddSystem(sList.ControlSys)
	w.AddSystem(sList.DragSys)
	w.AddSystem(sList.VelSys)
	w.AddSystem(sList.BoxSys)
	w.AddSystem(sList.LookSys)
	w.AddSystem(sList.CollisionSys)
	w.AddSystem(&HitSystem{NPlayers: ms.NPlayers})
	w.AddSystem(NewSpawnSystem(sList))

}

func AddBall(partner Draggable, pnum int, fric, l float32, sl SysList) *Ball {
	psc := partner.GetSpaceComponent().Position

	c := NewBall(psc.X+rand.Float32()*l/2, psc.Y+rand.Float32()*l/2, 10, pnum)
	sl.DragSys.Connect(partner, c, fric, l)
	sl.RenderSys.AddByInterface(c)
	sl.VelSys.Add(c)
	sl.CollisionSys.Add(c)
	return c
}
