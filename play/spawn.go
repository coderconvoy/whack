package play

import "engo.io/ecs"

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
		b := NewBoxy(-10, -10)
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
