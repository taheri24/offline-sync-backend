package internal

import (
	"github.com/common-nighthawk/go-figure"
)

func ShowBanner(text string) {
	myFigure := figure.NewFigure(text, "", true)
	myFigure.Print()
}
