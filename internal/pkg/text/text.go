package text

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"
)

var unicodeIcons = Icon{
	None: '·',
	HSep: '—',
	VSep: '∣',
	Size: '×',
	Pick: '⁞',
	Grep: '❯',
	Ps1:  '❯',
	Abr:  '…',
}

var defaultIcons = Icon{
	None: '.',
	HSep: '-',
	VSep: '|',
	Size: 'x',
	Pick: ':',
	Grep: '>',
	Ps1:  '>',
	Abr:  '…',
}

type Icon struct {
	None, HSep, VSep, Size, Pick, Grep, Ps1, Abr rune
}

func Int(s string) int {
	if n, err := strconv.Atoi(s); err != nil {
		return -1
	} else {
		return n
	}
}

func Dec(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func Len(s string) (l int) {
	return runewidth.StringWidth(s)
}

func Abr(s string, w int) string {
	if Len(s) > w {
		s = runewidth.TruncateLeft(s, Len(s)-w, "…")
	}

	return s
}

func Pad(s string, w int) string {
	return runewidth.FillRight(s, w)
}

func Trim(s string, l, r int) string {
	s = runewidth.TruncateLeft(s, l, "")
	s = runewidth.Truncate(s, r, "")

	return s
}

func Icons(u bool) *Icon {
	if u {
		return &unicodeIcons
	} else {
		return &defaultIcons
	}
}

func Title(l, r string, u bool) string {
	return fmt.Sprintf("%s %c %s", l, Icons(u).HSep, r)
}

func Block(s string, w int) (t string) {
	if w < 0 {
		w = 4 + len(s)
	}

	l := strings.Repeat("─", w-2)

	t += fmt.Sprintf("┌%s┐\n", l)
	t += fmt.Sprintf("│ %-*s │\n", w-4, s)
	t += fmt.Sprintf("└%s┘", l)

	return
}

func Human(n int64) string {
	const unit = 1000

	if n < unit {
		return fmt.Sprintf("%dB", n)
	}

	div, exp := int64(unit), 0

	for n := n / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f%c", float64(n)/float64(div), "KMGTPE"[exp])
}
