package gipfs

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/ipfs/pkg/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (g *GioIPFS) getPages() pages {
	return pages{
		"Welcome": gipfsPage{
			Title:  "Welcome",
			Header: g.welcomeHeader(),
			Body:   g.welcomeBody(),
		},
		"Status": gipfsPage{
			Title:  "Status",
			Header: g.statusHeader(),
			Body:   g.statusBody(),
		},
		"Files": gipfsPage{
			Title:  "Files",
			Header: g.filesHeader(),
			Body:   g.filesBody(),
		},
		"Explore": gipfsPage{
			Title:  "Explore",
			Header: g.exploreHeader(),
			Body:   g.exploreBody(),
		},
		"Peers": gipfsPage{
			Title:  "Peers",
			Header: g.peersHeader(),
			Body:   g.peersBody(),
		},
		"Settings": gipfsPage{
			Title:  "Settings",
			Header: g.settingsHeader(),
			Body:   g.settingsBody(),
		},
	}
}

func (g *GioIPFS) page(page gipfsPage) func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["White"], 0, 0, 0, 0, func(gtx C) D {
		return lyt.Format(gtx, "vflexs(start,r(inset(0dp30dp20dp30dp,_)),f(1,inset(0dp0dp0dp16dp,_)))",
			page.Header,
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				return contentList.Layout(gtx, len(page.Body), func(gtx C, i int) D {
					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
						page.Body[i],
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

func (g *GioIPFS) pageButton(b *widget.Clickable, f func(), icon, page string) func(gtx C) D {
	return func(gtx C) D {
		btn := material.IconButton(g.UI.Theme.T, b, g.UI.Theme.Icons[icon])
		btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
		btn.Size = unit.Dp(21)
		btn.Background = helper.HexARGB(g.UI.Theme.Colors["Secondary"])
		for b.Clicked() {
			f()
			currentPage = page
		}
		return btn.Layout(gtx)
	}
}
