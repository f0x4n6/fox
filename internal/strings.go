package fox

import (
	_ "embed"
)

const (
	Product = "Forensic Examiner"
	Website = "forensic-examiner.eu"
)

var (
	//go:embed res/banner.txt
	Banner string

	//go:embed res/version.txt
	Version string
)
