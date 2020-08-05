package gipfs

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"image"
)

func (g *GioIPFS) logo() func(gtx C) D {
	return func(gtx C) D {
		return material.Clickable(gtx, navBtn, func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			for navBtn.Clicked() {
				switch g.UI.mob {
				case false:
					g.UI.mob = true
				default:
					g.UI.mob = false
				}
			}
			inset := "inset(35dp20dp35dp45dp,_)"
			if g.UI.mob {
				inset = "inset(35dp20dp35dp28dp,_)"
			}
			return lyt.Format(gtx, inset, func(gtx C) D {
				logo := g.UI.logoText
				if g.UI.mob {
					logo = g.UI.logo
				}
				return logo.Layout(gtx)
			})
		})
	}
}

func toPointF(p image.Point) f32.Point {
	return f32.Point{X: float32(p.X), Y: float32(p.Y)}
}
