package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/lgylgy/gitw/gui"
)

var (
	sidebarView *gui.SidebarView
	branchView  *gui.BranchView
	remotesView *gui.RemotesView
	contentView *gui.ContentView
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
	g.SetManagerFunc(layout)

	// Side bar
	sidebarView = gui.NewSidebarView()
	if err != nil {
		log.Fatalln(err)
	}

	// Branch layout
	branchView = gui.NewBranchView()
	if err != nil {
		log.Fatalln(err)
	}

	// Remote layout
	remotesView = gui.NewRemotesView()
	if err != nil {
		log.Fatalln(err)
	}

	// Main content
	contentView = gui.NewContentView()
	if err != nil {
		log.Fatalln(err)
	}

	// Quit binding
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatalln(err)
	}

	// Main loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}

func layout(g *gocui.Gui) error {
	err := sidebarView.Draw(g)
	if err != nil {
		log.Fatalln(err)
	}
	err = branchView.Draw(g)
	if err != nil {
		log.Fatalln(err)
	}
	err = remotesView.Draw(g)
	if err != nil {
		log.Fatalln(err)
	}
	err = contentView.Draw(g)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
