package play

import "github.com/coderconvoy/engotil"

func RemoveBoxy(sl []*Boxy, b engotil.IDable) []*Boxy {
	id := b.ID()
	dp := -1
	for i, v := range sl {
		if v.ID() == id {
			dp = i
			break
		}
	}
	if dp >= 0 {
		return append(sl[:dp], sl[dp+1:]...)
	}
	return sl
}
