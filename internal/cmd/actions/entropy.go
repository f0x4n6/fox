package actions

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/internal"
	"github.com/cuhsat/fox/internal/opt/ui"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

var EntropyUsage = fox.Ascii + `
Display file entropy.

Usage:
  fox entropy [FLAG ...] PATH ...

Alias:
  en

Positional arguments:
  Path(s) to open

Global:
  -p, --print              print directly to console
  -h, --head               limit head of file by ...
  -t, --tail               limit tail of file by ...
  -n, --lines[=NUMBER]     number of lines (default: 10)
  -c, --bytes[=NUMBER]     number of bytes (default: 16)

Entropy:
      --min[=DECIMAL]      minimum score (default: 0.8)
      --max[=DECIMAL]      maximum score (default: 0.8)

Example:
  $ fox entropy -n ./**/*

Type "fox help" for more help...
`

var Entropy = &cobra.Command{
	Use:     "entropy",
	Short:   "display file entropy",
	Long:    "display file entropy",
	Aliases: []string{"en"},
	Args:    cobra.ArbitraryArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		flg := flags.Get()

		flg.Optional.NoConvert = true
		flg.Optional.NoPlugins = true

		if flg.Entropy.Min < 0 {
			sys.Exit("min must be 0 or greater")
		}

		if flg.Entropy.Max > 1 {
			sys.Exit("max must be 1 or lesser")
		}

		if flg.Entropy.Min > flg.Entropy.Max {
			sys.Exit("max must be greater than min")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print(EntropyUsage)
			os.Exit(2)
		} else if !flags.Get().Print {
			ui.Start(args, types.Entropy)
		} else {
			flg := flags.Get()

			hs := heapset.New(args)
			defer hs.ThrowAway()

			hs.Range(func(_ int, h *heap.Heap) bool {
				if v := h.Entropy(
					flg.Entropy.Min,
					flg.Entropy.Max,
				); v != -1 {
					fmt.Printf("%.10f  %s\n", v, h.String())
				}
				return true
			})
		}
	},
}

func init() {
	flg := flags.Get()

	Entropy.SetHelpTemplate(EntropyUsage)
	Entropy.Flags().BoolVarP(&flg.Print, "print", "p", false, "print directly to console")
	Entropy.Flags().BoolVarP(&flg.Limits.IsHead, "head", "h", false, "limit head of file by ...")
	Entropy.Flags().BoolVarP(&flg.Limits.IsTail, "tail", "t", false, "limit tail of file by ...")
	Entropy.Flags().IntVarP(&flg.Limits.Lines, "lines", "n", 0, "number of lines (default: 10)")
	Entropy.Flags().IntVarP(&flg.Limits.Bytes, "bytes", "c", 0, "number of bytes (default: 16)")
	Entropy.Flags().Float64Var(&flg.Entropy.Min, "min", 0.0, "minimum score")
	Entropy.Flags().Float64Var(&flg.Entropy.Max, "max", 1.0, "maximum score")
	Entropy.Flags().Lookup("min").NoOptDefVal = "0.8"
	Entropy.Flags().Lookup("max").NoOptDefVal = "0.8"
}
