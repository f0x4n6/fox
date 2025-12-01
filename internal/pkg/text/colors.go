package text

import "github.com/fatih/color"

var Crt = color.New(color.FgHiRed).SprintFunc()
var Err = color.New(color.FgRed).SprintFunc()
var Wrn = color.New(color.FgYellow).SprintFunc()

var Bold = color.New(color.Bold).SprintFunc()
var Hint = color.New(color.FgHiBlack).SprintFunc()
