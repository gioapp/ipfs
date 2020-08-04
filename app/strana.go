package gipfs

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/ipfs/pkg/itembtn"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"os"
)

var (
	contentList = &layout.List{
		Axis: layout.Vertical,
	}
)

func (w *GioIPFS) strana() func(gtx C) D {
	return w.Panel(w.Strana.Naziv, func(gtx C) D { return D{} }, w.itemsList(), func(gtx C) D { return D{} })
}

func (w *GioIPFS) listaAAA() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return contentList.Layout(gtx, len(w.Prikaz.e), func(gtx C, i int) D {
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				w.Prikaz.e[i],
				helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]),
			)
		})

	}
}

func (w *GioIPFS) itemsList() func(gtx C) D {
	return func(gtx C) D {
		return contentList.Layout(gtx, len(w.ItemsList), func(gtx C, i int) D {
			item := w.ItemsList[i]
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				func(gtx C) D {
					b := itembtn.ItemBtn(w.UI.Tema.T, item.btn, item.check, w.UI.Tema.Icons["GlyphFolder"], w.UI.Tema.Icons["GlyphDots"], item.Name, item.Hash, item.Size).Layout(gtx)
					for item.btn.Clicked() {
						w.lista("/" + item.Name)
					}
					return b
				},

				helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]),
			)
		})
	}
}

//func itemClicked(w *WingCMS,l sadrzaj.TipSadrzajaPrikaz) {
//	for l.Link.Clicked() {
//		w.Strana = WingStrana{l.Naziv, l.SlugMnozina}
//		//w.Prikaz = w.Db.DbReadAll(l.SlugMnozina)
//	}
//	return
//}

func (w *GioIPFS) lista(path string) {
	stat, err := w.sh.FilesStat(w.ctx, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	f, err := w.sh.ObjectGet(stat.Hash)
	checkError(err)

	w.ItemsList = generateItemsView(f.Links)
}

func generateItemsView(items []shell.ObjectLink) []folderListItem {
	var folder []folderListItem
	for _, item := range items {
		folder = append(folder, folderListItem{
			Name:  item.Name,
			Hash:  item.Hash,
			Size:  item.Size,
			btn:   new(widget.Clickable),
			check: new(widget.Bool),
		})
	}
	return folder
}
