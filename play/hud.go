package play

import (
	"fmt"
	"image/color"

	"github.com/coderconvoy/engotil"

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

type ScoreComponent struct {
	Health, Score int
	Hdisp, Sdisp  SetNumer
}

type HudSystem struct {
	RS     *common.RenderSystem
	fnt    *common.Font
	Scores []*ScoreComponent
}

func NewHudSystem(rs *common.RenderSystem, fnt *common.Font) *HudSystem {
	Scores := []*ScoreComponent{}

	return &HudSystem{
		RS:     rs,
		fnt:    fnt,
		Scores: Scores,
	}
}

func (hs *HudSystem) AddPlayer(pl *Boy) {
	i := len(hs.Scores) + 1
	hDisp := NewTextEntity("H:",
		common.SpaceComponent{Position: engo.Point{float32(i * 40), 0}},
		hs.fnt)
	hDisp.SetNum(5)
	sDisp := NewTextEntity("S:",
		common.SpaceComponent{Position: engo.Point{float32(i * 40), 40}},
		hs.fnt)
	hs.RS.AddByInterface(hDisp)
	hs.RS.AddByInterface(sDisp)

	res := &ScoreComponent{
		Health: 5,
		Score:  0,
		Hdisp:  hDisp,
		Sdisp:  sDisp,
	}
	hs.Scores = append(hs.Scores, res)
	pl.ScoreComponent = res

}

func (hs *HudSystem) New(w *ecs.World) {
	engo.Mailbox.Listen("GCollisionMessage", func(m engo.Message) {
		mes, ok := m.(engotil.GCollisionMessage)
		if !ok {
			return
		}

		if mes.Group&C_BOY_HURT > 0 {
			b, ok := mes.Main.(*Boy)
			if !ok {
				return
			}
			b.Health--
			b.Hdisp.SetNum(b.Health)
			if b.Health == 0 {
				fmt.Println("Death")
				engo.SetScene(&MainScene{NPlayers: len(hs.Scores)}, true)
			}
		}

	})
}

func (hs *HudSystem) Update(d float32) {

}

func (hs *HudSystem) Remove(o ecs.BasicEntity) {
}
