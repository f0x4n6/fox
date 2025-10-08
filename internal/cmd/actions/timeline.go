package actions

import (
	"cmp"
	"fmt"
	"os"
	"slices"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/internal"
	"github.com/cuhsat/fox/internal/opt/ui"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

var TimelineUsage = fox.Ascii + `
Display super timeline.

Usage:
  fox timeline [FLAG ...] PATH ...

Alias:
  super, tl

Positional arguments:
  Path(s) to open

Global:
  -p, --print              print directly to console

Timeline:
  -c, --cef                use Common Event Format

Example:
  $ fox timeline -c ./**/*.evtx

Type "fox help" for more help...
`

var Timeline = &cobra.Command{
	Use:     "timeline",
	Short:   "display super timeline",
	Long:    "display super timeline",
	Args:    cobra.ArbitraryArgs,
	Aliases: []string{"super", "tl"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print(TimelineUsage)
			os.Exit(2)
		} else if !flags.Get().Print {
			ui.Start(args, types.Timeline)
		} else {
			var ls []string

			flg := flags.Get()

			hs := heapset.New(args)
			defer hs.ThrowAway()

			hs.Range(func(_ int, h *heap.Heap) bool {
				for _, str := range *h.FMap() {
					if l := text.Normalize(str.Str, flg.Timeline.Cef); len(l) > 0 {
						ls = append(ls, l)
					}
				}
				return true
			})

			slices.SortStableFunc(ls, func(a, b string) int {
				return cmp.Compare(a, b)
			})

			for _, l := range ls {
				fmt.Println(l)
			}
		}
	},
}

func init() {
	flg := flags.Get()

	Timeline.SetHelpTemplate(TimelineUsage)
	Timeline.Flags().BoolVarP(&flg.Print, "print", "p", false, "print directly to console")
	Timeline.Flags().BoolVarP(&flg.Timeline.Cef, "cef", "c", false, "use Common Event Format")
}
