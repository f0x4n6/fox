package actions

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/v3/internal"
	"github.com/cuhsat/fox/v3/internal/opt/ui"
	"github.com/cuhsat/fox/v3/internal/pkg/flags"
	"github.com/cuhsat/fox/v3/internal/pkg/types"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heapset"
)

var CountsUsage = fox.Fox + `
Display line and byte counts.

Usage:
  fox counts [FLAG ...] PATH ...

Alias:
  wc

Positional arguments:
  Path(s) to open

Example:
  $ fox counts ./**/*.txt

Type "fox help" for more help...
`

var Counts = &cobra.Command{
	Use:     "counts",
	Short:   "display line and byte counts",
	Long:    "display line and byte counts",
	Args:    cobra.ArbitraryArgs,
	Aliases: []string{"wc"},
	PreRun: func(cmd *cobra.Command, args []string) {
		flg := flags.Get()

		flg.Optional.NoConvert = true
		flg.Optional.NoPlugins = true
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print(CountsUsage)
			os.Exit(2)
		} else if !flags.Get().Print {
			ui.Start(args, types.Counts)
		} else {
			hs := heapset.New(args)
			defer hs.ThrowAway()

			hs.Range(func(_ int, h *heap.Heap) bool {
				fmt.Printf("%8dL %8dB  %s\n", h.Length(), len(*h.MMap()), h.String())
				return true
			})
		}
	},
}

func init() {
	Counts.SetHelpTemplate(CountsUsage)
}
