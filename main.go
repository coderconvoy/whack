package main

import (
	"flag"

	"engo.io/engo"
	"github.com/coderconvoy/whack/play"
)

func main() {
	np := flag.Int("np", 1, "np: Number of players")
	flag.Parse()
	opts := engo.RunOptions{
		Title:         "Whack",
		ScaleOnResize: true,
		Width:         700,
		Height:        700,
	}
	engo.Run(opts, &play.MainScene{*np})
}
