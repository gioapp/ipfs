package gipfs

import (
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (g *GioIPFS) twoPanels(l, r int, left, right func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(start,f(0.60,inset(30dp10dp30dp10dp,_)),f(0.30,inset(30dp10dp30dp10dp,_)))",
			ContainerLayout(g.UI.Theme.Colors["PanelBg"], l, l, l, l, left),
			ContainerLayout(g.UI.Theme.Colors["PanelBg"], r, r, r, r, right))
	}
}
