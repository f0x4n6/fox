package text

import "github.com/fatih/color"

var Crit = color.New(color.FgHiRed).SprintFunc()
var High = color.New(color.FgRed).SprintFunc()
var Warn = color.New(color.FgYellow).SprintFunc()
var Info = color.New(color.FgWhite).SprintFunc()

var Bold = color.New(color.Bold).SprintFunc()
var Hint = color.New(color.FgHiBlack).SprintFunc()
