package fox

import (
	_ "embed"
)

const (
	Product = "Forensic Examiner"
	Website = "forensic-examiner.eu"
	Version = "v4.0.0"
)

var (
	//go:embed res/banner.txt
	Banner string
)
