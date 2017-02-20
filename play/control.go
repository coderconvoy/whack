package play

type ControlComponent struct {
	ctrls []KeyCommand
}

func (cc *ControlComponent) GetControls() []KeyCommand {
	return cc.ctrls
}

type KeyCommand struct {
	K    engo.Key
	S    string
	Push engo.Point
}

func GetKeys(n pnum) []KeyCommand {
	if n == 0 {
		return []KeyCommand{
			{engo.Left, "l1", engo.Point{-1, 0}},
			{engo.Right, "r1", engo.Point{1, 0}},
			{engo.Up, "u1", engo.Point{0, -1}},
			{engo.Down, "d1", engo.Point{0, 1}},
		}
	}
	return []KeyCommand{
		{engo.A, "l2", engo.Point{-1, 0}},
		{engo.D, "r2", engo.Point{1, 0}},
		{engo.W, "u2", engo.Point{0, -1}},
		{engo.S, "d2", engo.Point{0, 1}},
	}
}

type Controlable interface {
}
