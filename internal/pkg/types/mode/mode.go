package mode

import (
	"strings"
)

const (
	Default = Less
)

const (
	Less = Mode("Less")
	Grep = Mode("Grep")
	Pick = Mode("Pick")
	Goto = Mode("Goto")
	Open = Mode("Open")
	Chat = Mode("Chat")
	Hex  = Mode("Hex")
)

type Mode string

func (m Mode) String() string {
	return strings.ToUpper(string(m))
}

func (m Mode) IsFilter() bool {
	switch m {
	case Less, Grep, Pick:
		return true
	default:
		return false
	}
}

func (m Mode) IsStatic() bool {
	switch m {
	case Hex:
		return true
	default:
		return false
	}
}

func (m Mode) IsPrompt() bool {
	switch m {
	case Less, Hex:
		return false
	default:
		return true
	}
}

func (m Mode) IsSelect() bool {
	switch m {
	case Open:
		return true
	default:
		return false
	}
}
