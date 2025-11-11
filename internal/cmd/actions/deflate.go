package actions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cuhsat/fox/v3/internal"
	"github.com/cuhsat/fox/v3/internal/pkg/flags"
	"github.com/cuhsat/fox/v3/internal/pkg/sys"
	"github.com/cuhsat/fox/v3/internal/pkg/types"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v3/internal/pkg/types/heapset"
)

var DeflateUsage = fox.Fox + `
Deflate compressed files.

Usage:
  fox deflate [FLAG ...] PATH...

Alias:
  extract, unzip, de

Positional arguments:
  Path(s) to open

Additional flags:
  -l, --list               don't deflate, only list files
  -o, --out[=PATH]         output to path (default: .)

Example:
  $ fox deflate --pass=infected ioc.zip

Type "fox help" for more help...
`

var Deflate = &cobra.Command{
	Use:     "deflate",
	Short:   "deflate compressed files",
	Long:    "deflate compressed files",
	Args:    cobra.ArbitraryArgs,
	Aliases: []string{"extract", "unzip", "de"},
	PreRun: func(cmd *cobra.Command, args []string) {
		flg := flags.Get()

		flg.Optional.NoConvert = true
		flg.Optional.NoPlugins = true

		if flg.Optional.Readonly && !flg.Deflate.List {
			sys.Exit("deflate does write files")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print(DeflateUsage)
			os.Exit(2)
		} else if flags.Get().Deflate.List {
			hs := heapset.New(args)
			defer hs.ThrowAway()

			hs.Range(func(_ int, h *heap.Heap) bool {
				fmt.Println(strings.TrimPrefix(h.Title, h.Base)[1:])
				return true
			})
		} else {
			flg := flags.Get()

			hs := heapset.New(args)
			defer hs.ThrowAway()

			hs.Range(func(_ int, h *heap.Heap) bool {
				root := flg.Deflate.Path

				if root == "." {
					name := filepath.Base(h.Base)
					root = name[0 : len(name)-len(filepath.Ext(name))]
				}

				// convert to relative path
				path := h.Title

				if h.Type == types.Deflate {
					path = path[len(h.Base)+1:]
				} else {
					path = filepath.Base(path)
				}

				// create (sub)folders
				if sub := filepath.Dir(path); len(sub) > 0 {
					sub = filepath.Join(root, sub)

					err := os.MkdirAll(sub, 0700)

					if err != nil {
						sys.Exit(fmt.Sprintf("could not create path: %s", err.Error()))
					}
				}

				path = filepath.Join(root, path)

				// deflate file
				err := os.WriteFile(path, *h.MMap(), 0600)

				if err != nil {
					sys.Exit(fmt.Sprintf("could not deflate file: %s", err.Error()))
				}

				// calculate hash
				if !flg.NoFile {
					sum, err := h.HashSum(types.SHA256)

					if err != nil {
						sys.Exit(fmt.Sprintf("could not compute hash: %s", err.Error()))
					}

					fmt.Printf("%x  %s\n", sum, path)
				}
				return true
			})

			fmt.Printf("%d file(s) written\n", hs.Len())
		}
	},
}

func init() {
	flg := flags.Get()

	Deflate.SetHelpTemplate(DeflateUsage)
	Deflate.Flags().BoolVarP(&flg.Deflate.List, "list", "l", false, "don't deflate, only list files")
	Deflate.Flags().StringVarP(&flg.Deflate.Path, "out", "o", "", "output to path")
	Deflate.Flags().Lookup("out").NoOptDefVal = "."
	_ = Deflate.MarkFlagDirname("out")
}
