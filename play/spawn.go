package play

import (
	"math/rand"

	"engo.io/ecs"
)

type SpawnSystem struct {
	S     SysList
	Delay float32
	since float32
}

func NewSpawnSystem(sl SysList) *SpawnSystem {
	return &SpawnSystem{
		S:     sl,
		Delay: 3,
	}
}

func (ss *SpawnSystem) Update(d float32) {
	ss.since += d
	if ss.since > ss.Delay {
		var b *Boxy
		switch rand.Intn(4) {
		case 0:
			b = NewBoxy(-10, 700)
		case 1:
			b = NewBoxy(-10, -10)
		case 2:
			b = NewBoxy(700, 700)
		default:
			b = NewBoxy(700, -10)

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
