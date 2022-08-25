package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/git"
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

 /gitw config.json

`)
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) != 2 {
		log.Print("missing configuration file")
		usage()
	}

	// Config file
	config, err := git.LoadConfiguration(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()

	// Create views
	layout := NewLayout(g, config)
	g.SetManagerFunc(func(g *gocui.Gui) error {
		return layout.Draw(g)
	})

	// Key binding
	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
		log.Fatalln(err)
	}

	// Main loop
	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
