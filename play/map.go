package play

import (
	"fmt"

	"github.com/coderconvoy/engotil"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

func LoadMap(fname string, sl SysList) *common.Level {
	reso, err := engo.Files.Resource(fname)
	if err != nil {
		panic(err)
	}
	tmxR := reso.(common.TMXResource)
	levelData := tmxR.Level
	fmt.Println("Level Data Loaded")

	//var tiles []*Tile

	for _, tLayer := range levelData.TileLayers {
		for _, tElem := range tLayer.Tiles {
			if tElem.Image != nil {
				tile := &Tile{
					BasicEntity: ecs.NewBasic(),
					RenderComponent: common.RenderComponent{
						Drawable: tElem,
						Scale:    engo.Point{0.7, 0.7},
					},
					SpaceComponent: common.SpaceComponent{
						Position: engo.Point{tElem.X * 0.7, tElem.Y * 0.7},
						Width:    tElem.Width() * 0.7,
						Height:   tElem.Height() * 0.7,
					},
					GCollisionComponent: engotil.GCollisionComponent{
						Main:  0,
						Extra: engo.Point{0, 0},
					},
				}
				sl.RenderSys.AddByInterface(tile)
				if tLayer.Name == "sea" {
					tile.Group = C_BOY_SOLID
					sl.CollisionSys.Add(tile)
				}
				if tLayer.Name == "ground" {
					tile.Group = C_BOY_SOLID | C_MOVING_SOLID
					sl.CollisionSys.Add(tile)
				}
			}
		}
	}

	return levelData
}

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	engotil.GCollisionComponent
}
