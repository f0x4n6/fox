package actions

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/internal"
	"github.com/cuhsat/fox/internal/opt/ui"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/internal/pkg/types/page"
)

var StringsUsage = fox.Fox + `
Display ASCII and Unicode strings.

Usage:
  fox strings [FLAG ...] PATH ...

Alias:
  carve, st

Positional arguments:
  Path(s) to open

Additional flags:
      --min=NUMBER         minimum length (default: 3)
      --max=NUMBER         maximum length (default: Unlimited)
      --ascii              carve only ASCII strings
      --class              run built-in classification:
						     ipv4, ipv6, mac, mail, url, uuid

Example:
  $ fox strings --class malware.exe

Type "fox help" for more help...
`

var Strings = &cobra.Command{
	Use:     "strings",
	Short:   "display ASCII and Unicode strings",
	Long:    "display ASCII and Unicode strings",
	Args:    cobra.ArbitraryArgs,
	Aliases: []string{"carve", "st"},
	PreRun: func(cmd *cobra.Command, args []string) {
		flg := flags.Get()

		flg.Optional.NoConvert = true
		flg.Optional.NoPlugins = true

		re, err := cmd.Flags().GetString("regex")

		if len(re) > 0 {
			flg.Strings.Re, err = regexp.Compile(re)

			if err != nil {
				sys.Exit(fmt.Sprintf("could not compile regex: %s", err.Error()))
			}
		}

		if flg.Strings.Min <= 0 {
			sys.Exit("min must be greater than 0")
		}

		if flg.Strings.Max <= 0 {
			sys.Exit("min must be greater than 0")
		}

		if flg.Strings.Min > flg.Strings.Max {
			sys.Exit("max must be greater than min")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print(StringsUsage)
			os.Exit(2)
		} else if !flags.Get().Print {
			ui.Start(args, types.Strings)
		} else {
			flg := flags.Get()

			hs := heapset.New(args)
			defer hs.ThrowAway()

			hs.Range(func(_ int, h *heap.Heap) bool {
				if h.Type != types.Stdin {
					if !flg.NoFile {
						fmt.Println(text.Block(h.String(), page.TermW))
					}

					for s := range h.Strings(
						flg.Strings.Min,
						flg.Strings.Max,
						flg.Strings.Class,
						flg.Strings.Re,
					) {
						if !flg.NoLine {
							fmt.Printf("%08x  %s\n", s.Off, strings.TrimSpace(s.Str))
						} else {
							fmt.Println(strings.TrimSpace(s.Str))
						}
					}
				}
				return true
			})
		}
	},
}

func init() {
	flg := flags.Get()

	Strings.SetHelpTemplate(StringsUsage)
	Strings.Flags().IntVar(&flg.Strings.Min, "min", 3, "minimum length")
	Strings.Flags().IntVar(&flg.Strings.Max, "max", math.MaxInt, "maximum length")
	Strings.Flags().BoolVar(&flg.Strings.Ascii, "ascii", false, "carve only ASCII strings")
	Strings.Flags().BoolVar(&flg.Strings.Class, "class", false, "run built-in classification")
}
