package text

import (
	"fmt"
	"math"
	"strings"
)

func Dec(n int) int {
	return int(math.Log10(float64(n))) + 1
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
