package actions

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/v3/internal"
	"github.com/cuhsat/fox/v3/internal/opt/ui"
	"github.com/cuhsat/fox/v3/internal/pkg/flags"
	"github.com/cuhsat/fox/v3/internal/pkg/sys"
	"github.com/cuhsat/fox/v3/internal/pkg/text"
	"github.com/cuhsat/fox/v3/internal/pkg/types"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heapset"
	"github.com/cuhsat/fox/v3/internal/pkg/types/page"
)

var HashUsage = fox.Fox + `
Display file hash or checksums.

Usage:
  fox hash [FLAG ...] PATH ...

Alias:
  ha

Positional arguments:
  Path(s) to open

Additional flags:
      --type=ALGORITHM     hash algorithm (default: SHA256)

    Cryptographic hash algorithms:
      MD5, SHA1, SHA256, SHA3, SHA3-224, SHA3-256, SHA3-384, SHA3-512

    Performance hash algorithms:
      FNV-1, FNV-1A, XXH64, XXH3

    Similarity hash algorithms:
      SDHASH, SSDEEP, TLSH

    Checksum algorithms:
      CRC32-IEEE, CRC64-ECMA, CRC64-ISO

Example:
  $ fox hash --type=md5 --type=sha1 artifacts.zip

Type "fox help" for more help...
`

var Hash = &cobra.Command{
	Use:     "hash",
	Short:   "display file hash or checksums",
	Long:    "display file hash or checksums",
	Args:    cobra.ArbitraryArgs,
	Aliases: []string{"ha"},
	PreRun: func(cmd *cobra.Command, args []string) {
		flg := flags.Get()

		flg.Optional.NoConvert = true
		flg.Optional.NoPlugins = true

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
			ui.Start(args, types.HashSum)
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
	Hash.Flags().Var(&flg.Hash.Algos, "type", "hash algorithm")
}
