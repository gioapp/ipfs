package gipfs

import (
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/ipfs/pkg/theme"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"math"
	"strconv"
)

func (g *GioIPFS) GetStatus() {
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
	return
}

func (g *GioIPFS) GetLiveStat() {

	fbw, err := g.sh.StatsBW(g.ctx)
	checkError(err)
	live = &statLive{
		//HostingSize: "uint",
		RateOut:  fbw.RateOut,
		RateIn:   fbw.RateIn,
		TotalIn:  fbw.TotalIn,
		TotalOut: fbw.TotalOut,
	}
}

func (g *GioIPFS) statusHeader() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["PanelBg"], 10, 10, 10, 10, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		helper.Fill(gtx, helper.HexARGB(g.UI.Theme.Colors["PanelBg"]))
		return lyt.Format(gtx, "vflexb(middle,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp30dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)))",
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H1(g.UI.Theme, "Connected to IPFS")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.UI.Theme, "Hosting 54.1 MB of files — Discovered 149 peers")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			statusRow(g.UI.Theme, "PEER ID", row(g.UI.Theme, g.Status.PeerId)),
			statusRow(g.UI.Theme, "VERSION", row(g.UI.Theme, g.Status.Version)),
			statusRow(g.UI.Theme, "GATEWAY", row(g.UI.Theme, g.Status.Gateway)),
			statusRow(g.UI.Theme, "API", row(g.UI.Theme, g.Status.Api)),
			statusRow(g.UI.Theme, "ADDRESSES", ContainerLayout(g.UI.Theme.Colors["White"], 10, 10, 10, 10, func(gtx C) D {
				return addressesList.Layout(gtx, len(g.Status.Addresses), func(gtx C, i int) D {
					title := theme.Body(g.UI.Theme, g.Status.Addresses[i])
					title.Alignment = text.Start
					return title.Layout(gtx)
				})
			})),
			statusRow(g.UI.Theme, "PUBLIC KEY", ContainerLayout(g.UI.Theme.Colors["White"], 10, 10, 10, 10, func(gtx C) D {
				title := theme.Body(g.UI.Theme, g.Status.Pub)
				title.Alignment = text.Start
				return title.Layout(gtx)
			})))
	})
}

func row(th *theme.Theme, label string) func(gtx C) D {
	return func(gtx C) D {
		t := theme.Body(th, label)
		t.Alignment = text.Start
		return t.Layout(gtx)
	}
}

func statusRow(th *theme.Theme, label string, content func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(start,r(inset(0dp16dp0dp0dp,_)),f(1,_))",
			//gtx.Constraints.Min.X = gtx.Constraints.Max.X
			func(gtx C) D {
				gtx.Constraints.Min.X = 100
				title := theme.Body(th, label)
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			content,
		)
	}
}

func (g *GioIPFS) statusBody() []func(gtx C) D {
	return []func(gtx C) D{
		func(gtx C) D {
			var (
				rateIn   string = "0"
				rateOut  string = "0"
				totalIn  string = "0"
				totalOut string = "0"
			)
			//if live != nil {
			if live.RateIn != 0 {
				rateIn = formatByteSize(live.RateIn)
			}
			if live.RateOut != 0 {
				rateOut = formatByteSize(live.RateOut)
			}
			if live.TotalIn != 0 {
				totalIn = formatByteSize(float64(live.TotalIn))
			}
			if live.TotalOut != 0 {
				totalOut = formatByteSize(float64(live.TotalOut))
			}

			//fmt.Println("gore", formatByteSize(0))
			//}
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_),r(_),r(_))",
				statusRow(g.UI.Theme, "RateIn: ", row(g.UI.Theme, rateIn)),
				statusRow(g.UI.Theme, "RateOut: ", row(g.UI.Theme, rateOut)),
				statusRow(g.UI.Theme, "TotalIn: ", row(g.UI.Theme, totalIn)),
				statusRow(g.UI.Theme, "TotalOut: ", row(g.UI.Theme, totalOut)),
			)
			//return D{}
		},
	}
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func formatByteSize(size float64) string {
	//size := sizeInMB * 1024 * 1024
	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	//fmt.Println("dole", strconv.FormatFloat(getSize, 'f', -1, 64)+" "+string(getSuffix))
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}
