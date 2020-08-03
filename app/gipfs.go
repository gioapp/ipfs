package gipfs

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/ipfs/pkg/nav"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

type Page struct {
	layout func(layout.Context) layout.Dimensions
	//nav.NavItem
	//Actions  []materials.AppBarAction
	//Overflow []materials.OverflowAction
}

func (w *GioIPFS) GlavniEkran(gtx layout.Context) {
	lyt.Format(gtx, "hflexb(start,r(_),f(1,_))",
		//w.Meni(),
		func(gtx C) D {
			n := nav.Navigation{
				Name:  "Navigacion",
				Bg:    w.UI.Tema.Colors["Primary"],
				Items: menuItems,
			}
			return n.Nav(w.UI.Tema, gtx)
		},
		func(gtx C) D {
			//return lyt.Format(gtx, "vflexb(start,r(_),f(1,_),r(_))",
			return lyt.Format(gtx, "vflexb(start,f(1,_))",
				//header(w),
				w.strana(),
				//footer(w),
			)
		})
}

func (w *GioIPFS) cell(align text.Alignment, tekst string) func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		cell := material.Caption(w.UI.Tema.T, tekst)
		cell.TextSize = unit.Dp(12)
		cell.Alignment = align
		return cell.Layout(gtx)
	}
}

func (w *GioIPFS) sumaFooter(t string) func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		suma := material.Body2(w.UI.Tema.T, t)
		suma.Alignment = text.End
		return suma.Layout(gtx)
	}
}
