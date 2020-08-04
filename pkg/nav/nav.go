package nav

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/theme"
	"github.com/gioapp/ipfs/pkg/icontextbtn"
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

func (n *Navigation) Nav(th *theme.DuoUItheme, gtx layout.Context) layout.Dimensions {
	//var nav []func(gtx layout.Context) layout.Dimensions
	//nav = append(nav, func(gtx layout.Context) layout.Dimensions {
	//	return material.Body1(th.T, "testr").Layout(gtx)
	//})
	gtx.Constraints.Max.X = 180
	return navList.Layout(gtx, len(n.Items), func(gtx layout.Context, i int) layout.Dimensions {
		item := n.Items[i]

		btn := icontextbtn.IconTextBtn(th.T, item.Btn, item.Icon, unit.Dp(48), th.Colors["White"], item.Title)

		btn.IconSize = unit.Dp(72)
		btn.TextSize = unit.Dp(12)
		btn.Axis = layout.Horizontal
		btn.CornerRadius = unit.Dp(0)
		btn.Background = helper.HexARGB(th.Colors["Primary"])
		return btn.Layout(gtx)
	})
}
