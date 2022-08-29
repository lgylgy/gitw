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
	repositories := git.NewRepositories()
	err := repositories.Load(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	// Create views
	layout := NewLayout(g, repositories)
	g.SetManagerFunc(func(g *gocui.Gui) error {
		layout.Manage()
		return nil
	})

	// Key binding
	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
		log.Fatalln(err)
	}

	// Main loop
	layout.Manage()
	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
