package main

import (
	"fmt"
	"gioui.org/app"
	_ "gioui.org/app/permission/storage"
	"gioui.org/io/system"
	"gioui.org/layout"
	gipfs "github.com/gioapp/ipfs/app"
	"github.com/gioapp/ipfs/cfg"
	in "github.com/gioapp/ipfs/cfg/ini"
	"log"
	"os"
	"time"
)

func main() {

	g := gipfs.NewGioIPFS()

	if cfg.Initial {
		fmt.Println("running initial sync")
	}
	in.Init(g.Settings.File)
	ticker(g.Tik())

	go func() {
		defer os.Exit(0)
		if err := loop(g); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(g *gipfs.GioIPFS) error {
	for {
		select {
		case e := <-g.UI.Window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				g.UI.Context = layout.NewContext(&g.UI.Ops, e)

				g.BeforeMain()
				//if !g.API.OK {
				//g.GreskaEkran()
				//} else {
				g.AppMain()
				//}
				g.AfterMain()

				e.Frame(g.UI.Context.Ops)
			}
			g.UI.Window.Invalidate()
		}
	}
}

func ticker(f func()) {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				f()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
