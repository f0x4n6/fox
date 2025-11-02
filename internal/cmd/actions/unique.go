package actions

import (
	"fmt"
	"os"

	"github.com/cuhsat/fox/internal/pkg/types/page"
	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/internal"
	"github.com/cuhsat/fox/internal/opt/ui"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

var UniqueUsage = fox.Fox + `
Display unique lines.

Usage:
  fox unique [FLAG ...] PATH ...

Alias:
  anew, un

Positional arguments:
  Path(s) to open

Example:
  $ fox unique ./**/*.log

Type "fox help" for more help...
`

var Unique = &cobra.Command{
	Use:     "unique",
	Short:   "display unique lines",
	Long:    "display unique lines",
	Args:    cobra.ArbitraryArgs,
	Aliases: []string{"anew", "un"},
	Run: func(cmd *cobra.Command, args []string) {
		flg := flags.Get()

		if len(args) == 0 {
			fmt.Print(UniqueUsage)
			os.Exit(2)
		} else if !flg.Print {
			ui.Start(args, types.Unique)
		} else {
			hs := heapset.New(args)
			defer hs.ThrowAway()

			hs.Unique().CloseOther()

			buf := page.NewContext(hs.LoadHeap())

			for l := range page.Text(buf).Lines {
				fmt.Println(l)
			}
		}
	},
}

func init() {
	Unique.SetHelpTemplate(UniqueUsage)
}
