package play

import (
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type SpawnSystem struct {
	S     SysList
	Delay float32
	since float32
	rats  []*common.Spritesheet
}

func NewSpawnSystem(sl SysList) *SpawnSystem {
	return &SpawnSystem{
		S:     sl,
		Delay: 3,
		rats: []*common.Spritesheet{
			common.NewSpritesheetFromFile("rat.png", 40, 40),
			common.NewSpritesheetFromFile("rat2.png", 40, 40),
		},
	}
}

func (ss *SpawnSystem) Update(d float32) {
	ss.since += d
	n := rand.Intn(2)
	if ss.since > ss.Delay {
		var b *Boxy
		switch rand.Intn(4) {
		case 0:
			b = NewBoxy(n, -10, 700, ss.rats[n])
		case 1:
			b = NewBoxy(n, -10, -10, ss.rats[n])
		case 2:
			b = NewBoxy(n, 700, 700, ss.rats[n])
		default:
			b = NewBoxy(n, 700, -10, ss.rats[n])

		}

		ss.S.BoxSys.AddBox(b)
		ss.S.CollisionSys.Add(b)
		ss.S.RenderSys.AddByInterface(b)
		ss.S.VelSys.Add(b)

		ss.since = 0
		ss.Delay -= 0.01
	}
}
func (ss *SpawnSystem) Remove(e ecs.BasicEntity) {
}
