package actions

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/v3/internal"
	"github.com/cuhsat/fox/v3/internal/opt/ui"
	"github.com/cuhsat/fox/v3/internal/pkg/flags"
	"github.com/cuhsat/fox/v3/internal/pkg/types"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heapset"
)

var TimelineUsage = fox.Fox + `
Display super timeline.

Usage:
  fox timeline [FLAG ...] PATH ...

Alias:
  time, tl

Positional arguments:
  Path(s) to open

Additional flags:
      --cef                use the Common Event Format

Example:
  $ fox timeline --cef ./**/*.evtx

Type "fox help" for more help...
`

var Timeline = &cobra.Command{
	Use:     "timeline",
	Short:   "display super timeline",
	Long:    "display super timeline",
	Args:    cobra.ArbitraryArgs,
	Aliases: []string{"time", "tl"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print(TimelineUsage)
			os.Exit(2)
		} else if !flags.Get().Print {
			ui.Start(args, types.Timeline)
		} else {
			cef := flags.Get().Timeline.Cef

			hs := heapset.New(args)
			defer hs.ThrowAway()

			for _, l := range hs.Extract(cef) {
				fmt.Println(l)
			}
		}
	},
}

func init() {
	flg := flags.Get()

	Timeline.SetHelpTemplate(TimelineUsage)
	Timeline.Flags().BoolVar(&flg.Timeline.Cef, "cef", false, "use the Common Event Format")
}
