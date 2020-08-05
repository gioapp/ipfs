package gipfs

import (
	"context"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"github.com/gioapp/ipfs/pkg/theme"
	shell "github.com/ipfs/go-ipfs-api"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	selected int
)

type GioIPFS struct {
	sh        *shell.Shell
	ctx       context.Context
	UI        gipfsUI
	menuItems []Item
	Settings  gipfsSettings
	ItemsList []*folderListItem
	Status    Status
}

type folderListItem struct {
	Name  string
	Hash  string
	Size  uint64
	Type  uint8
	btn   *widget.Clickable
	check *widget.Bool
}

type gipfsUI struct {
	Device  string
	Window  *app.Window
	Theme   *theme.Theme
	Context layout.Context
	//Ekran   func(gtx layout.Context) layout.Dimensions
	FontSize float32

	mob   bool
	pages pages
	Ops   op.Ops
}

type gipfsSettings struct {
	Dir  string
	File string
}

type gipfsPage struct {
	Title  string
	Header func(gtx C) D
	Body   []func(gtx C) D
}

type pages map[string]gipfsPage

type Navigation struct {
	Name  string
	Bg    string
	Logo  Logo
	Items []Item
}
type Item struct {
	Title string
	Bg    string
	Icon  *widget.Icon
	Btn   *widget.Clickable
	//Page *gipfsPage
}

type Logo struct {
	Title string
	Logo  string
}

type Status struct {
	Title       string
	HostingSize uint
	PeerId      string
	Version     string
	Gateway     string
	Api         string
	Addresses   []string
	Pub         string
}
type statLive struct {
	RateOut  float64
	RateIn   float64
	TotalIn  int64
	TotalOut int64
}
