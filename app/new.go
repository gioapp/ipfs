package gipfs

import (
	"context"
	"fmt"
	"gioui.org/app"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/ipfs/pkg/icon/icons"
	"github.com/gioapp/ipfs/pkg/nav"
	"github.com/gioapp/ipfs/pkg/theme"
	shell "github.com/ipfs/go-ipfs-api"
	"image"
	"image/png"
	"os"
)

var (
	ipfsLogoTextImageOp = paint.ImageOp{}
	ipfsLogoImageOp     = paint.ImageOp{}
	ipfsLogoTextImage   image.Image
	ipfsLogoImage       image.Image
)

func NewGioIPFS() *GioIPFS {

	g := &GioIPFS{
		//Db:     jdb.New("db"),
		sh:              shell.NewShell("/ip4/127.0.0.1/tcp/5001"),
		ctx:             context.Background(),
		daemonConnected: make(chan bool),
	}

	g.UI = gipfsUI{
		Theme: theme.NewTheme(),
		//mob:   make(chan bool),
	}
	g.UI.Theme.Icons = icons.NewIPFSicons()
	g.UI.Theme.T.Color.Primary = helper.HexARGB(g.UI.Theme.Colors["Primary"])
	g.UI.Theme.T.Color.Text = helper.HexARGB(g.UI.Theme.Colors["Charcoal"])
	g.UI.Theme.T.Color.Hint = helper.HexARGB(g.UI.Theme.Colors["Silver"])
	g.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1280), unit.Dp(1024)),
		app.Title("IPFS"),
	)
	g.menuItems = []nav.Item{
		nav.Item{
			Title: "Status",
			Icon:  g.UI.Theme.Icons["StrokeMarketing"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Files",
			Icon:  g.UI.Theme.Icons["StrokeWeb"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Explore",
			Icon:  g.UI.Theme.Icons["StrokeIpld"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Peers",
			Icon:  g.UI.Theme.Icons["StrokeCube"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Settings",
			Icon:  g.UI.Theme.Icons["StrokeSettings"],
			Btn:   new(widget.Clickable),
		},
	}

	g.GetStatus()
	g.Status.Live = statLive{}

	g.Page = gipfsPage{
		Title:  "Status",
		Header: g.statusHeader(),
		Body:   g.statusBody(),
	}

	ipfsLogoTextImageFile, err := os.Open("/home/marcetin/go/src/github.com/gioapp/ipfs/pkg/icon/logo/ipfs-logo-text.png")
	checkError(err)
	defer ipfsLogoTextImageFile.Close()
	ipfsLogoTextImage, err = png.Decode(ipfsLogoTextImageFile)
	checkError(err)

	ipfsLogoImageFile, err := os.Open("/home/marcetin/go/src/github.com/gioapp/ipfs/pkg/icon/logo/ipfs-logo.png")
	checkError(err)
	defer ipfsLogoImageFile.Close()
	ipfsLogoImage, err = png.Decode(ipfsLogoImageFile)
	checkError(err)

	logoText := paint.NewImageOp(ipfsLogoTextImage)
	g.UI.logoText = widget.Image{
		Src:   logoText,
		Scale: 1,
	}
	logo := paint.NewImageOp(ipfsLogoImage)
	g.UI.logo = widget.Image{
		Src:   logo,
		Scale: 1,
	}
	return g
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
