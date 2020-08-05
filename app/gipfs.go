package gipfs

import (
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (g *GioIPFS) AppMain() {
	lyt.Format(g.UI.Context, "hflexb(start,r(_),f(1,_))",
		func(gtx C) D {
			gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
			n := Navigation{
				Name:  "Navigacion",
				Bg:    g.UI.Theme.Colors["NavBg"],
				Items: g.menuItems,
			}
			width := 252
			if g.UI.mob {
				width = 128
			}
			return n.Nav(g.UI.Theme, gtx, width, g.UI.mob, g.logo())
		},
		func(gtx C) D {
			return lyt.Format(gtx, "vflexb(start,r(_),f(1,_))",
				g.header(),
				g.page(g.UI.pages[currentPage]),
			)
		})
}

func (g *GioIPFS) BeforeMain() {

}

func (g *GioIPFS) AfterMain() {

}

func (g *GioIPFS) Tik() func() {
	return func() {
		g.GetLiveStat()
	}
}
