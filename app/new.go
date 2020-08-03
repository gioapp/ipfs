package gipfs

import (
	"context"
	"fmt"
	"gioui.org/app"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/theme"
	"github.com/gioapp/ipfs/pkg/icon/icons"
	"github.com/gioapp/ipfs/pkg/nav"
	shell "github.com/ipfs/go-ipfs-api"
)

func NewGioIPFS() *GioIPFS {

	w := &GioIPFS{
		Strana: gipfsStrana{"Komandna tabla", "komandna_tabla"},
		//Db:     jdb.New("db"),
		sh:  shell.NewShell("localhost:5001"),
		ctx: context.Background(),
	}

	w.UI = gipfsUI{
		Tema: theme.NewDuoUItheme(),
	}
	w.UI.Tema.Icons = icons.NewIPFSicons()

	w.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1280), unit.Dp(1024)),
		app.Title(w.Podesavanja.Naziv),
	)

	menuItems = []nav.Item{
		nav.Item{
			Title: "Status",
			Icon:  th.Icons["StrokeMarketing"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Files",
			Icon:  "StrokeWeb",
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Explore",
			Icon:  "StrokeIpld",
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Peers",
			Icon:  "StrokeCube",
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Settings",
			Icon:  "StrokeSettings",
			Btn:   new(widget.Clickable),
		},
	}

	return w
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
