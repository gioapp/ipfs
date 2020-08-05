package gipfs

import (
	"context"
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/ipfs/pkg/itembtn"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

var (
	contentList = &layout.List{
		Axis: layout.Vertical,
	}
	addressesList = &layout.List{
		Axis: layout.Vertical,
	}
)

func (g *GioIPFS) itemsList() func(gtx C) D {
	return func(gtx C) D {
		return contentList.Layout(gtx, len(g.ItemsList), func(gtx C, i int) D {
			item := g.ItemsList[i]
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				func(gtx C) D {
					b := itembtn.ItemBtn(g.UI.Theme, item.btn, item.check, g.UI.Theme.Icons["GlyphFolder"], g.UI.Theme.Icons["GlyphDots"], item.Name, item.Hash, item.Size).Layout(gtx)
					for item.btn.Clicked() {
						fmt.Println("Name", "/"+item.Name)

						g.ItemsList = listFolder(g.ctx, g.sh, "/"+item.Name)
					}
					return b
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
