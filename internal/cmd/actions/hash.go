package actions

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/internal"
	"github.com/cuhsat/fox/internal/app/ui"
	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/internal/pkg/types/page"
)

var HashUsage = fox.Ascii + `
Display file hash or checksums.

Usage:
  fox hash [FLAG ...] PATH ...

Alias:
  sum, ha

Positional arguments:
  Path(s) to open

Global:
  -p, --print              print directly to console

Hash:
  -t, --type=ALGORITHM     hash algorithm (default: SHA256)

    Cryptographic hash algorithms:
      MD5, SHA1, SHA256, SHA3, SHA3-224, SHA3-256, SHA3-384, SHA3-512

    Similarity hash algorithms:
      SDHASH, SSDEEP, TLSH

    Performance hash algorithms:
      XXH64, XXH3

    Checksum algorithms:
      CRC32-IEEE, CRC64-ECMA, CRC64-ISO

Example:
  $ fox hash -t=md5 -t=sha1 artifacts.zip

Type "fox help" for more help...
`

var Hash = &cobra.Command{
	Use:     "hash",
	Short:   "display file hash or checksums",
	Long:    "display file hash or checksums",
	Args:    cobra.ArbitraryArgs,
	Aliases: []string{"sum", "ha"},
	PreRun: func(cmd *cobra.Command, args []string) {
		flg := flags.Get()

		flg.Opt.NoConvert = true
		flg.Opt.NoPlugins = true

		if flg.Print {
			flg.Opt.NoConvert = true
			flg.Opt.NoPlugins = true
		}

		// default
		if len(flg.Hash.Algos.Value) == 0 {
			_ = flg.Hash.Algos.Set(types.SHA256)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print(HashUsage)
			os.Exit(2)
		} else if !flags.Get().Print {
			ui.Start(args, types.Hash)
		} else {
			algos := flags.Get().Hash.Algos.Value

			hs := heapset.New(args)
			defer hs.ThrowAway()

			for _, algo := range algos {
				if len(algos) > 1 {
					fmt.Println(text.Block(strings.ToUpper(algo), page.TermW))
				}

				hs.Range(func(_ int, h *heap.Heap) bool {
					sum, err := h.HashSum(algo)

					if err != nil {
						sys.Exit(fmt.Sprintf("could not compute hash: %s", err.Error()))
						return false
					}

					switch algo {
					case types.SDHASH:
						fmt.Printf("%s  %s\n", sum, h.String())
					default:
						fmt.Printf("%x  %s\n", sum, h.String())
					}
					return true
				})
			}
		}
	},
}

func init() {
	flg := flags.Get()

	Hash.SetHelpTemplate(HashUsage)
	Hash.Flags().BoolVarP(&flg.Print, "print", "p", false, "print directly to console")
	Hash.Flags().VarP(&flg.Hash.Algos, "type", "t", "hash algorithm")
}
