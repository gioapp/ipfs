package gipfs

import (
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"image"
)

var (
	submitBtn         = new(widget.Clickable)
	tourBtn           = new(widget.Clickable)
	welcomeBtn        = new(widget.Clickable)
	navBtn            = new(widget.Clickable)
	browseBtn         = new(widget.Clickable)
	headerSearchInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	apiAddressInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	navList = &layout.List{
		Axis: layout.Vertical,
	}
	contentList = &layout.List{
		Axis: layout.Vertical,
	}
	addressesList = &layout.List{
		Axis: layout.Vertical,
	}
	ipfsLogoTextImageOp = paint.ImageOp{}
	ipfsLogoTextImage   image.Image
	ipfsLogoImageOp     = paint.ImageOp{}
	ipfsLogoImage       image.Image
	ipldLogoImageOp     = paint.ImageOp{}
	ipldLogoImage       image.Image
	currentPage         string
	daemonConnected     bool

	logoTextImage widget.Image
	logoImage     widget.Image
	logoIpldImage widget.Image

	live = &statLive{}

	suffixes = [7]string{
		0: "B",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
		5: "PB",
		6: "EB",
	}
)
