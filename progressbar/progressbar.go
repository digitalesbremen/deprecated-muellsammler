package progressbar

import (
	"fmt"

	"github.com/schollz/progressbar/v2"
)

type ProgressBar struct {
	progressbar progressbar.ProgressBar
}

func (p *ProgressBar) Add(num int) {
	_ = p.progressbar.Add(num)
}

func (p *ProgressBar) Finish(format string, a ...interface{}) {
	_ = p.progressbar.Finish()
	fmt.Println()
	fmt.Printf(format, a)
	fmt.Println()
	fmt.Println()
}

func NewProgressBar(max int) *ProgressBar {
	bar := progressbar.NewOptions(max, progressbar.OptionSetRenderBlankState(true))
	return &ProgressBar{*bar}
}
