package gipfs

import (
	"gioui.org/text"
	"github.com/gioapp/ipfs/pkg/theme"
)

func (g *GioIPFS) GetExplore() {
	f, err := g.sh.ID()
	checkError(err)
	fv, ft, err := g.sh.Version()
	checkError(err)
	g.Status = Status{
		Title: "string",
		//HostingSize: "uint",
		PeerId:  f.ID,
		Version: fv + " " + ft,
		//Gateway:   ,
		//Api:       "string",
		Addresses: f.Addresses,
		Pub:       f.PublicKey,
	}
}

func (g *GioIPFS) exploreHeader() func(gtx C) D {
	return func(gtx C) D {
		//return ContainerLayout(g.UI.Theme.Colors["PanelBg"], 10, 10, 10, 10, func(gtx C) D {
		//
		//})
		return D{}
	}
}

func (g *GioIPFS) exploreBody() []func(gtx C) D {
	return []func(gtx C) D{
		//func(gtx C) D {
		//	gtx.Constraints.Min.X = gtx.Constraints.Max.X
		//	return lyt.Format(gtx, "hflexb(middle,f(0.6,inset(5dp0dp5dp0dp,_)),f(0.4, inset(5dp0dp5dp0dp,_))))",
		//		//helper.Fill(gtx, helper.HexARGB(g.UI.Theme.Colors["PanelBg"]))
		//		func(gtx C) D {
		//			return lyt.Format(gtx, "vflexb(middle,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)))",
		//				func(gtx C) D {
		//					gtx.Constraints.Min.X = gtx.Constraints.Max.X
		//					title := theme.H1(g.UI.Theme, "Explore the Merkle Forest\n")
		//					title.Alignment = text.Start
		//					return title.Layout(gtx)
		//				},
		//				func(gtx C) D {
		//					gtx.Constraints.Min.X = gtx.Constraints.Max.X
		//					title := theme.H6(g.UI.Theme, "Paste a CID into the box to fetch the IPLD node it addresses, or choose a featured dataset.\n\n")
		//					title.Alignment = text.Start
		//					return title.Layout(gtx)
		//				},
		//				statusRow(g.UI.Theme, "PUBLIC KEY", ContainerLayout(g.UI.Theme.Colors["White"], 10, 10, 10, 10, func(gtx C) D {
		//					title := theme.Body(g.UI.Theme, g.Status.Pub)
		//					title.Alignment = text.Start
		//					return title.Layout(gtx)
		//				})))
		//		},
		//		ContainerLayout(g.UI.Theme.Colors["Silver"], 1, 1, 1, 1, ContainerLayout(g.UI.Theme.Colors["PanelBg"], 30, 30, 30, 30, func(gtx C) D {
		//			gtx.Constraints.Min.X = gtx.Constraints.Max.X
		//			return lyt.Format(gtx, "vflexb(middle,f(0.6,inset(5dp0dp5dp0dp,_)),f(0.4, inset(5dp0dp5dp0dp,_))))",
		//				func(gtx C) D {
		//					gtx.Constraints.Max.Y = gtx.Constraints.Min.Y
		//					return g.6UI.logoIpld.Layout(gtx)
		//				},
		//				func(gtx C) D {
		//					title := theme.Body(g.UI.Theme, "IPLD is the data model of the content-addressable web. It allows us to treat all hash-linked data structures as subsets of a unified information space, unifying all data models that link data with hashes as instances of IPLD.\n\nContent addressing through hashes has become a widely-used means of connecting data in distributed systems, from the blockchains that run your favorite cryptocurrencies, to the commits that back your code, to the web’s content at large. Yet, whilst all of these tools rely on some common primitives, their specific underlying data structures are not interoperable.\n\nEnter IPLD: a single namespace for all hash-inspired protocols. Through IPLD, links can be traversed across protocols, allowing you to explore data regardless of the underlying protocol.")
		//					title.Alignment = text.Start
		//					return title.Layout(gtx)
		//				})
		//		})))
		func(gtx C) D {
			title := theme.Body(g.UI.Theme, "LIVE DATA")
			title.Alignment = text.Start
			return title.Layout(gtx)
		},
	}
}
