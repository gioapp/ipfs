package gipfs

import (
	"context"
	"fmt"
	"gioui.org/app"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/ipfs/pkg/helper"
	"github.com/gioapp/ipfs/pkg/icon/icons"
	"github.com/gioapp/ipfs/pkg/theme"
	shell "github.com/ipfs/go-ipfs-api"
	"image/png"
	"os"
)

func NewGioIPFS() *GioIPFS {
	g := &GioIPFS{
		//Db:     jdb.New("db"),
		//sh:  shell.NewShell("/ip4/127.0.0.1/tcp/5011"),
		ctx: context.Background(),
	}

	sh := shell.NewShell("/ip4/127.0.0.1/tcp/5001")
	if sh != nil {
		fmt.Println("shshshshshPRE", sh)
		g.sh = sh
		fmt.Println("shshshshsPOSLEh", sh)
	}
	pwd = append(pwd, "Home")
	fmt.Println("ggggsh", g.sh)

	g.UI = gipfsUI{
		Theme: theme.NewTheme(),
		//mob:   make(chan bool),

	}
	currentPage = "Status"
	g.UI.Theme.Icons = icons.NewIPFSicons()
	g.UI.pages = g.getPages()
	g.UI.Theme.T.Color.Primary = helper.HexARGB(g.UI.Theme.Colors["Primary"])
	g.UI.Theme.T.Color.Text = helper.HexARGB(g.UI.Theme.Colors["Charcoal"])
	g.UI.Theme.T.Color.Hint = helper.HexARGB(g.UI.Theme.Colors["Silver"])
	g.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1024), unit.Dp(800)),
		app.Title("IPFS"),
	)
	g.menuItems = g.getMenuItems()
	//for _, item :=range g.menuItems{
	//	for item.Btn.Clicked(){
	//		g.UI.currentPage = item.Title
	//		fmt.Println("ttt",item.Title)
	//
	//	}
	//}
	//g.Status.Live = statLive{}
	//g.Page = gipfsPage{
	//	Title:  "Status",
	//	Header: g.statusHeader(),
	//	Body:   g.statusBody(),
	//}
	getImages()
	g.GetStatus()
	g.GetFiles()

	//suffixes[0] = "B"
	//suffixes[1] = "KB"
	//suffixes[2] = "MB"
	//suffixes[3] = "GB"
	//suffixes[4] = "TB"

	return g
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}

func (g *GioIPFS) getMenuItems() []Item {
	return []Item{
		Item{
			Title: "Status",
			Icon:  g.UI.Theme.Icons["StrokeMarketing"],
			Btn:   new(widget.Clickable),
		},
		Item{
			Title: "Files",
			Icon:  g.UI.Theme.Icons["StrokeWeb"],
			Btn:   new(widget.Clickable),
		},
		Item{
			Title: "Explore",
			Icon:  g.UI.Theme.Icons["StrokeIpld"],
			Btn:   new(widget.Clickable),
		},
		Item{
			Title: "Peers",
			Icon:  g.UI.Theme.Icons["StrokeCube"],
			Btn:   new(widget.Clickable),
		},
		Item{
			Title: "Settings",
			Icon:  g.UI.Theme.Icons["StrokeSettings"],
			Btn:   new(widget.Clickable),
		},
	}
}

func getImages() {
	ipfsLogoTextImageFile, err := os.Open("./pkg/icon/logo/ipfs-text.png")
	checkError(err)
	defer ipfsLogoTextImageFile.Close()
	ipfsLogoTextImage, err = png.Decode(ipfsLogoTextImageFile)
	checkError(err)

	ipfsLogoImageFile, err := os.Open("./pkg/icon/logo/ipfs.png")
	checkError(err)
	defer ipfsLogoImageFile.Close()
	ipfsLogoImage, err = png.Decode(ipfsLogoImageFile)
	checkError(err)

	ipldLogoImageFile, err := os.Open("./pkg/icon/logo/ipld.png")
	checkError(err)
	defer ipldLogoImageFile.Close()
	ipldLogoImage, err = png.Decode(ipldLogoImageFile)
	checkError(err)

	logoText := paint.NewImageOp(ipfsLogoTextImage)
	logoTextImage = widget.Image{
		Src:   logoText,
		Scale: 1,
	}
	logo := paint.NewImageOp(ipfsLogoImage)
	logoImage = widget.Image{
		Src:   logo,
		Scale: 1,
	}
	logoIpld := paint.NewImageOp(ipldLogoImage)
	logoIpldImage = widget.Image{
		Src:   logoIpld,
		Scale: 1,
	}
}
