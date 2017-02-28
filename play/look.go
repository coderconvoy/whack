package play

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"github.com/coderconvoy/engotil"
	"github.com/coderconvoy/engotil/engopoint"
)

type LookFace interface {
	common.BasicFace
	common.SpaceFace
	LookAngle(n int)
}

type LookEntity struct {
	a LookFace
	b common.SpaceFace
}

type LookSystem struct {
	obs []LookEntity
}

func (ls *LookSystem) Connect(a LookFace, b common.SpaceFace) {
	ls.obs = append(ls.obs, LookEntity{a, b})
}
func (ls *LookSystem) Update(d float32) {
	for _, i := range ls.obs {
		asc := i.a.GetSpaceComponent()
		acen := engotil.SpaceCenter(*asc)
		bsc := i.b.GetSpaceComponent()
		bcen := engotil.SpaceCenter(*bsc)
		vec := engopoint.Sub(bcen, acen)
		a := engopoint.Angle8(vec)
		i.a.LookAngle(a)
	}
}
func (ls *LookSystem) Remove(e ecs.BasicEntity) {
}
