package gipfs

import (
	"context"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"github.com/gioapp/ipfs/pkg/nav"
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
	Page gipfsPage
	//Db                   *jdb.JavazacDB
	sh  *shell.Shell
	ctx context.Context
	UI  gipfsUI

	menuItems []nav.Item
	Settings  gipfsSettings

	ItemsList []*folderListItem

	Status Status

	daemonConnected chan bool
}

//type prikaz struct {
//	w map[string]interface{}
//	e []func(gtx C) D
//}

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
	logoText widget.Image
	logo     widget.Image
	mob      bool
	Ops      op.Ops
}

type gipfsSettings struct {
	Dir  string
	File string
}
