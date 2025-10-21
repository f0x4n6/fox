package actions

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/internal"
	"github.com/cuhsat/fox/internal/opt/ui"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

var CountsUsage = fox.Ascii + `
Display line and byte counts.

Usage:
  fox counts [FLAG ...] PATH ...

Alias:
  co, wc

Positional arguments:
  Path(s) to open

Global:
  -p, --print              print directly to console
  -h, --head               limit head of file by ...
  -t, --tail               limit tail of file by ...
  -n, --lines[=NUMBER]     number of lines (default: 10)
  -c, --bytes[=NUMBER]     number of bytes (default: 16)

Example:
  $ fox counts ./**/*.txt

Type "fox help" for more help...
`

var Counts = &cobra.Command{
	Use:     "counts",
	Short:   "display line and byte counts",
	Long:    "display line and byte counts",
	Args:    cobra.ArbitraryArgs,
	Aliases: []string{"co", "wc"},
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
	flg := flags.Get()

	Counts.SetHelpTemplate(CountsUsage)
	Counts.Flags().BoolVarP(&flg.Print, "print", "p", false, "print directly to console")
	Counts.Flags().BoolVarP(&flg.Limits.IsHead, "head", "h", false, "limit head of file by ...")
	Counts.Flags().BoolVarP(&flg.Limits.IsTail, "tail", "t", false, "limit tail of file by ...")
	Counts.Flags().IntVarP(&flg.Limits.Lines, "lines", "n", 0, "number of lines (default: 10)")
	Counts.Flags().IntVarP(&flg.Limits.Bytes, "bytes", "c", 0, "number of bytes (default: 16)")
}
