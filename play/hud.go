package play

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type TextEntity struct {
	Pre string
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	fnt *common.Font
}

func NewTextEntity(pre string, sc common.SpaceComponent, fnt *common.Font) *TextEntity {
	db := common.Text{
		Font: fnt,
		Text: pre + "0",
	}

	res := &TextEntity{
		Pre:         pre,
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Drawable: db,
			Color:    color.Black,
		},
		SpaceComponent: sc,
		fnt:            fnt,
	}

	res.SetZIndex(5)
	res.SetShader(common.HUDShader)
	return res

}

func (te *TextEntity) SetNum(n int) {
	te.RenderComponent.Drawable = common.Text{
		Font: te.fnt,
		Text: fmt.Sprintf("%s%d", te.Pre, n),
	}

}

type SetNumer interface {
	common.BasicFace
	SetNum(int)
}

type Player struct {
	Health, Score int
	Hdisp, Sdisp  SetNumer
}

type HudSystem struct {
	RS      *common.RenderSystem
	Players []*Player
}

func NewHudSystem(np int, rs *common.RenderSystem, fnt *common.Font) *HudSystem {

	players := []*Player{}
	for i := 0; i < np; i++ {
		hDisp := NewTextEntity("H:",
			common.SpaceComponent{Position: engo.Point{float32(i * 40), 0}},
			fnt)
		sDisp := NewTextEntity("S:",
			common.SpaceComponent{Position: engo.Point{float32(i * 40), 40}},
			fnt)
		rs.AddByInterface(hDisp)
		rs.AddByInterface(sDisp)

		players = append(players, &Player{
			Health: 5,
			Score:  0,
			Hdisp:  hDisp,
			Sdisp:  sDisp,
		})
	}

	return &HudSystem{
		RS:      rs,
		Players: players,
	}

}

func (hs *HudSystem) New(w *ecs.World) {
}

func (hs *HudSystem) Update(d float32) {
	for _, p := range hs.Players {
		p.Score++
		p.Hdisp.SetNum(p.Score)
	}

}

func (hs *HudSystem) Remove(o ecs.BasicEntity) {
}
