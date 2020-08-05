package gipfs

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

var (
	headerSearchInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
)

func (g *GioIPFS) header() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["Info"], 0, 0, 0, 0, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "hflexb(middle,r(inset(0dp0dp0dp6dp,_)),r(inset(20dp30dp20dp3dp,_)))",
			g.headerSearch(),
			g.headerMenu(),
		)
	})
}
func (g *GioIPFS) headerMenu() func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_))",
			g.pageButton(welcomeBtn, func() {}, "StrokeCase"),
			helper.DuoUIline(true, 0, 2, 2, g.UI.Theme.Colors["DarkGrayI"]),
			g.pageButton(tourBtn, func() {}, "GlyphPencil"),
		)
	}

}

func (g *GioIPFS) headerSearch() func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,r(inset(20dp0dp20dp30dp,_)),r(_))",
			ContainerLayout(g.UI.Theme.Colors["White"], 8, 8, 8, 8, func(gtx C) D {
				gtx.Constraints.Min.X = 430
				e := material.Editor(g.UI.Theme.T, headerSearchInput, "QmHash")
				return e.Layout(gtx)
			}),
			func(gtx C) D {
				e := material.Button(g.UI.Theme.T, welcomeBtn, "Browse")
				e.Inset = layout.Inset{
					Top:    unit.Dp(4),
					Right:  unit.Dp(4),
					Bottom: unit.Dp(4),
					Left:   unit.Dp(4),
				}
				e.CornerRadius = unit.Dp(4)
				return e.Layout(gtx)
			},
		)
	}
}
