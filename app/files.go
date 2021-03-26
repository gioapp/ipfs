package gipfs

import (
	"context"
	"fmt"
	"gioui.org/widget/material"
	"github.com/ipfs/go-cid"

	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/ipfs/pkg/helper"
	"github.com/gioapp/ipfs/pkg/itembtn"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

var (
	l = &layout.List{
		Axis: layout.Vertical,
	}
	upBtn    = new(widget.Clickable)
	checkAll = new(widget.Bool)
)

type I []*FolderListItem

type FolderListItem struct {
	Name  string
	Cid   cid.Cid
	Size  uint64
	Type  uint8
	Btn   *widget.Clickable
	Check *widget.Bool
}

func (g *GioIPFS) itemsList() func(gtx C) D {

	//pwd = append(pwd, "")
	//fmt.Println("pwd", pwd)
	fmt.Println("NamePAPAICA", pathGen(pwd))
	list, err := g.sh.FilesLs(g.ctx, pathGen(pwd))
	//checkError(err)
	//list, err := g.sh.FilesLs(g.ctx, "/")
	checkError(err)
	itms := makeList(list)
	return func(gtx layout.Context) layout.Dimensions {
		return lyt.Format(gtx, "vflexb(middle,r(_),f(1,_))",
			func(gtx layout.Context) layout.Dimensions {
				return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
					func(gtx layout.Context) layout.Dimensions {
						b := itembtn.ItemBtn(g.UI.Theme, upBtn, checkAll, g.UI.Theme.Icons["GlyphFolder"], g.UI.Theme.Icons["GlyphDots"], "..", "", 0).Layout(gtx)
						for upBtn.Clicked() {
							//parentCid = c

							fmt.Println("Name", "/")

							items, err := g.sh.FilesLs(g.ctx, "/")
							checkError(err)
							itms = makeList(items)
						}
						return b
					},
					helper.DuoUIline(false, 0, 0, 1, g.UI.Theme.Colors["Gray"]),
				)
			},
			func(gtx layout.Context) layout.Dimensions {
				return l.Layout(gtx, len(itms), func(gtx layout.Context, i int) layout.Dimensions {
					itm := itms[i]
					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
						func(gtx layout.Context) layout.Dimensions {
							b := itembtn.ItemBtn(g.UI.Theme, itm.Btn, itm.Check, g.UI.Theme.Icons["GlyphFolder"], g.UI.Theme.Icons["GlyphDots"], itm.Name, itm.Cid.String(), itm.Size).Layout(gtx)
							for itm.Btn.Clicked() {
								//parentCid = c
								pwd = append(pwd, itm.Name)
								fmt.Println("pwd", pwd)

								fmt.Println("Name", pathGen(pwd))

								items, err := g.sh.FilesLs(g.ctx, pathGen(pwd))
								checkError(err)
								itms = makeList(items) //g.ItemsList = listFolder(g.ctx, g.sh, "/"+itm.Name)
								//itms = g.jdb.ReadList( "QmSv66pvzJfjwLHuQCYhd3cekGWNX6Q2o5Y268SNMw8fd8")
							}
							return b
						},
						helper.DuoUIline(false, 0, 0, 1, g.UI.Theme.Colors["Gray"]),
					)
				})
			},
		)
	}
}

func (g *GioIPFS) filesHeader() func(gtx C) D {
	return func(gtx C) D {
		return pwdList.Layout(gtx, len(pwd), func(gtx C, i int) D {
			item := pwd[i]
			return lyt.Format(gtx, "hflexb(middle,r(_),r(_))",
				func(gtx C) D {
					//b := itembtn.ItemBtn(g.UI.Theme, item.btn, item.check, g.UI.Theme.Icons["GlyphFolder"], g.UI.Theme.Icons["GlyphDots"], item.Name, item.Hash, item.Size).Layout(gtx)
					//for item.btn.Clicked() {
					//	fmt.Println("Name", "/"+item.Name)
					//
					//	g.ItemsList = listFolder(g.ctx, g.sh, "/"+item.Name)
					//}
					//return b
					return material.H6(g.UI.Theme.T, item).Layout(gtx)
				},
				helper.DuoUIline(false, 0, 0, 1, g.UI.Theme.Colors["Gray"]),
			)
		})
	}
}

func listFolder(ctx context.Context, sh *shell.Shell, path string) []*folderListItem {
	f, err := sh.FilesLs(ctx, path, shell.FilesLs.Stat(true))
	checkError(err)
	var folder []*folderListItem
	for _, item := range f {
		fmt.Println("item", item)
		folder = append(folder, &folderListItem{
			Name:  item.Name,
			Hash:  item.Hash,
			Size:  item.Size,
			Type:  item.Type,
			btn:   new(widget.Clickable),
			check: new(widget.Bool),
		})
	}
	return folder
}
func (g *GioIPFS) filesBody() []func(gtx C) D {
	return []func(gtx C) D{
		g.itemsList(),
		//items.ItemsList(g.g.UI.Theme, "QmUn3oue7CxL3ERQH26P8wQMZBmdhxrapHukg2AJwwBGEK")
		//return D{}
	}
}

func (g *GioIPFS) GetFiles() {
	items, err := g.sh.FilesLs(g.ctx, "/")
	checkError(err)

	for _, iii := range items {

		fmt.Println("ffffffffffffffffffffffff", iii.Name)
	}

	//makeList(items)
	//g.ItemsList = ffffff
	return
}
