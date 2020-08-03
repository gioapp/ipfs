package gipfs

import (
	"context"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"github.com/gioapp/gel/theme"
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
	Strana gipfsStrana
	//Db                   *jdb.JavazacDB
	sh  *shell.Shell
	ctx context.Context
	UI  gipfsUI

	Podesavanja gipfsPodesavanja
	Prikaz      prikaz
}

type prikaz struct {
	w map[string]interface{}
	e []func(gtx C) D
}

type folderListItem struct {
	Name  string
	Hash  string
	Size  uint64
	btn   *widget.Clickable
	check *widget.Bool
}

type gipfsUI struct {
	Device  string
	Window  *app.Window
	Tema    *theme.DuoUItheme
	Context layout.Context
	Ekran   func(gtx layout.Context) layout.Dimensions
	Ops     op.Ops
}

type gipfsPodesavanja struct {
	Naziv string
	Dir   string
	File  string
	Cyr   bool
}

type gipfsStrana struct {
	Naziv string
	Slug  string
}
