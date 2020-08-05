package nav

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/ipfs/pkg/icontextbtn"
	"github.com/gioapp/ipfs/pkg/theme"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

var (
	navList = &layout.List{
		Axis: layout.Vertical,
	}
)

type Navigation struct {
	Name  string
	Bg    string
	Logo  Logo
	Items []Item
}
type Item struct {
	Title string
	Bg    string
	Icon  *widget.Icon
	Btn   *widget.Clickable
}

type Logo struct {
	Title string
	Logo  string
}

func (n *Navigation) Nav(th *theme.Theme, gtx layout.Context, width int, noText bool, logo func(gtx layout.Context) layout.Dimensions) layout.Dimensions {
	gtx.Constraints.Min.X = width
	gtx.Constraints.Max.X = width
	gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
	helper.Fill(gtx, helper.HexARGB(n.Bg))
	return lyt.Format(gtx, "vflexs(start,r(_),f(1,_))",
		logo,
		func(gtx layout.Context) layout.Dimensions {
			return navList.Layout(gtx, len(n.Items), func(gtx layout.Context, i int) layout.Dimensions {
				item := n.Items[i]
				btn := icontextbtn.IconTextBtn(th, item.Btn, item.Icon, unit.Dp(85), th.Colors["Info"], item.Title, noText)
				btn.TextSize = unit.Dp(16)
				btn.CornerRadius = unit.Dp(0)
				btn.Background = helper.HexARGB(th.Colors["NavBg"])
				return btn.Layout(gtx)
			})
		})
}
