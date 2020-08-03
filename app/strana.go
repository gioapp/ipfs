package gipfs

import (
	"context"
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/theme"
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
	return w.Panel(w.Strana.Naziv, func(gtx C) D { return D{} }, lista(w.ctx, w.sh, w.UI.Tema, "/"), func(gtx C) D { return D{} })
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

func lista(ctx context.Context, sh *shell.Shell, th *theme.DuoUItheme, path string) func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		stat, err := sh.FilesStat(ctx, path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
		f, err := sh.ObjectGet(stat.Hash)
		checkError(err)

		files := generateItemsView(f.Links)

		return contentList.Layout(gtx, len(files), func(gtx C, i int) D {
			file := files[i]
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				func(gtx C) D {
					return lyt.Format(gtx, "hflexb(middle,r(_),r(_),f(0.8,_),r(_),f(0.2,_),r(_))",
						func(gtx C) D {
							return material.CheckBox(th.T, file.check, "").Layout(gtx)
						},
						func(gtx C) D {
							return material.Body1(th.T, "0").Layout(gtx)
						},
						func(gtx C) D {
							return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
								func(gtx C) D {
									return material.Body1(th.T, file.Name).Layout(gtx)
								},
								func(gtx C) D {
									return material.Body1(th.T, file.Hash).Layout(gtx)
								},
							)
						},
						func(gtx C) D {
							return material.Body1(th.T, "0").Layout(gtx)
						},
						func(gtx C) D {
							return material.Body1(th.T, fmt.Sprint(file.Size)).Layout(gtx)
						},
						func(gtx C) D {
							return material.Body1(th.T, "0").Layout(gtx)
						},
					)
				},
				helper.DuoUIline(false, 0, 0, 1, th.Colors["Gray"]),
			)
		})
	}
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
