package gipfs

import (
	"context"
	"fmt"
	"github.com/ipfs/go-cid"

	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
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
func (g *GioIPFS) filesBody() []func(gtx C) D {
	return []func(gtx C) D{
		ItemsList(g.ctx, g.sh, g.UI.Theme, "QmUn3oue7CxL3ERQH26P8wQMZBmdhxrapHukg2AJwwBGEK"),
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
