package actions

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/v3/internal"
	"github.com/cuhsat/fox/v3/internal/opt/ui"
	"github.com/cuhsat/fox/v3/internal/pkg/flags"
	"github.com/cuhsat/fox/v3/internal/pkg/sys"
	"github.com/cuhsat/fox/v3/internal/pkg/types"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heapset"
)

var EntropyUsage = fox.Fox + `
Display file entropy.

Usage:
  fox entropy [FLAG ...] PATH ...

Alias:
  en

Positional arguments:
  Path(s) to open

Additional flags:
      --min[=DECIMAL]      minimum score (default: 0.8)
      --max[=DECIMAL]      maximum score (default: 0.8)

Example:
  $ fox entropy --min ./**/*

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
	Entropy.Flags().Float64Var(&flg.Entropy.Min, "min", 0.0, "minimum score")
	Entropy.Flags().Float64Var(&flg.Entropy.Max, "max", 1.0, "maximum score")
	Entropy.Flags().Lookup("min").NoOptDefVal = "0.8"
	Entropy.Flags().Lookup("max").NoOptDefVal = "0.8"
}
