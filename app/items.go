package gipfs

import (
	"fmt"
	_ "gioui.org/app/permission/storage"
	"gioui.org/layout"
	"gioui.org/widget"
	shell "github.com/ipfs/go-ipfs-api"
)

func makeList(item []*shell.MfsLsEntry) I {
	var itms I
	for _, item := range item {
		itms = append(itms, &FolderListItem{
			Name: item.Name,
			//Cid:  item.Cid,
			Size: item.Size,
			//Type:  uint8,
			Btn:   new(widget.Clickable),
			Check: new(widget.Bool),
		})
		fmt.Println("NameMakeLISt", item.Name)

		//fmt.Println("CidCidMakeLISt", item.Cid.)
	}
	return itms
}

//func ItemsList(ctx context.Context, sh *shell.Shell, th *theme.Theme, hash string) func(gtx layout.Context) layout.Dimensions {
//	//c, _ := cid.Decode(hash)
//	//var parentCid cid.Cid
//	list, err := sh.FilesLs(ctx, "/")
//	checkError(err)
//	itms := makeList(list)
//	return func(gtx layout.Context) layout.Dimensions {
//		return lyt.Format(gtx, "vflexb(middle,r(_),f(1,_))",
//			func(gtx layout.Context) layout.Dimensions {
//				return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
//					func(gtx layout.Context) layout.Dimensions {
//						b := itembtn.ItemBtn(th, upBtn, checkAll, th.Icons["GlyphFolder"], th.Icons["GlyphDots"], "..", "", 0).Layout(gtx)
//						for upBtn.Clicked() {
//							//parentCid = c
//
//							fmt.Println("Name", "/")
//
//							items, err := sh.FilesLs(ctx, "/")
//							checkError(err)
//							itms = makeList(items)
//						}
//						return b
//					},
//					helper.DuoUIline(false, 0, 0, 1, th.Colors["Gray"]),
//				)
//			},
//			func(gtx layout.Context) layout.Dimensions {
//				return l.Layout(gtx, len(itms), func(gtx layout.Context, i int) layout.Dimensions {
//					itm := itms[i]
//					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
//						func(gtx layout.Context) layout.Dimensions {
//							b := itembtn.ItemBtn(th, itm.Btn, itm.Check, th.Icons["GlyphFolder"], th.Icons["GlyphDots"], itm.Name, itm.Cid.String(), itm.Size).Layout(gtx)
//							for itm.Btn.Clicked() {
//								//parentCid = c
//								g.pwd = append(g.pwd, itm.Name)
//								fmt.Println("pwd", g.pwd)
//
//								fmt.Println("Name",  pathGen(g.pwd))
//
//								items, err := sh.FilesLs(ctx, pathGen(g.pwd))
//								checkError(err)
//								itms = makeList(items) //g.ItemsList = listFolder(g.ctx, g.sh, "/"+itm.Name)
//								//itms = g.jdb.ReadList( "QmSv66pvzJfjwLHuQCYhd3cekGWNX6Q2o5Y268SNMw8fd8")
//							}
//							return b
//						},
//						helper.DuoUIline(false, 0, 0, 1, th.Colors["Gray"]),
//					)
//				})
//			},
//		)
//	}
//}

func pathGen(path []string) string {
	var p string
	for _, folder := range path {
		if folder != "Home" {
			p = p + "/" + folder
		}
	}
	return p
}

//
//func listFolder(ctx context.Context, path string) I {
//	//f, err := sh.FilesLs(ctx, path, shell.FilesLs.Stat(true))
//	//checkError(err)
//	var folder []*FolderListItem
//	for _, item := range f {
//		fmt.Println("item", item)
//		folder = append(folder, &FolderListItem{
//			Name:  item.Name,
//			//Cid:  item.Hash,
//			Size:  item.Size,
//			Type:  item.Type,
//			Btn:   new(widget.Clickable),
//			Check: new(widget.Bool),
//		})
//	}
//	return folder
//}

func filesBody() []func(gtx layout.Context) layout.Dimensions {
	return []func(gtx layout.Context) layout.Dimensions{
		func(gtx layout.Context) layout.Dimensions {
			//return lyt.Format(gtx, "vflexb(middle,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp30dp0dp,_)),r(inset(5dp0dp5dp0dp,_),r(inset(5dp0dp5dp0dp,_)))",
			//	statusRow(g.UI.Theme, "RateIn: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.RateIn))),
			//	statusRow(g.UI.Theme, "RateOut: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.RateOut))),
			//	statusRow(g.UI.Theme, "TotalIn: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.TotalIn))),
			//	statusRow(g.UI.Theme, "TotalOut: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.TotalOut))),
			//)
			return layout.Dimensions{}
		},
	}
}
