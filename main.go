package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

func usage() {
	fmt.Fprint(flag.CommandLine.Output(), `
    **** LGY Git GUI ****

 --- Repo ---  ----- Content ----
 |          |  |                |
 |          |  |                |
 ------------  ------------------
 --- Branches --- --- Remotes ---
 |              | |             |
 ---------------- ---------------
 --- Log-------------------------
 |                             |
 --------------------------------
`)
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()

	// Create views
	layout := NewLayout(g)
	g.SetManagerFunc(func(g *gocui.Gui) error {
		return layout.Draw(g)
	})

	// Key binding
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatalln(err)
	}

	// Main loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
