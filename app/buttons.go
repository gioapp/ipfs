package gipfs

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
)

var (
	tourBtn    = new(widget.Clickable)
	welcomeBtn = new(widget.Clickable)
	navBtn     = new(widget.Clickable)
)

func (g *GioIPFS) pageButton(b *widget.Clickable, f func(), icon string) func(gtx C) D {
	return func(gtx C) D {
		btn := material.IconButton(g.UI.Theme.T, b, g.UI.Theme.Icons[icon])
		btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
		btn.Size = unit.Dp(21)
		btn.Background = helper.HexARGB(g.UI.Theme.Colors["Secondary"])
		for b.Clicked() {
			f()
			g.Page = gipfsPage{}
		}
		return btn.Layout(gtx)
	}
}
