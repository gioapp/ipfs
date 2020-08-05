package gipfs

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

type gipfsPage struct {
	Title  string
	Header func(gtx C) D
	Body   []func(gtx C) D
}

func (g *GioIPFS) page() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["White"], 0, 0, 0, 0, func(gtx C) D {
		return lyt.Format(gtx, "vflexs(start,r(inset(0dp30dp20dp30dp,_)),f(1,inset(0dp0dp0dp16dp,_)))",
			g.Page.Header,
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				return contentList.Layout(gtx, len(g.Page.Body), func(gtx C, i int) D {
					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
						g.Page.Body[i],
						helper.DuoUIline(false, 0, 0, 1, g.UI.Theme.Colors["Gray"]),
					)
				})
			})
	})
}

func ContainerLayout(bg string, paddingTop, paddingRight, paddingBottom, paddingLeft int, itemContent func(gtx C) D) func(gtx C) D {
	//vmin := gtx.Constraints.Min.Y
	//if d.FullWidth {
	//	hmin = gtx.Constraints.Max.Y
	//}
	return func(gtx C) D {
		return layout.Stack{Alignment: layout.W}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				return helper.Fill(gtx, helper.HexARGB(bg))
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				//gtx.Constraints.Min = hmin
				//gtx.Constraints.Min = vmin
				return layout.Inset{
					Top:    unit.Dp(float32(paddingTop)),
					Right:  unit.Dp(float32(paddingRight)),
					Bottom: unit.Dp(float32(paddingBottom)),
					Left:   unit.Dp(float32(paddingLeft)),
				}.Layout(gtx, itemContent)
			}),
		)
	}
}
