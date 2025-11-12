package flags

import (
	"regexp"
)

type Flags struct {
	Bag bool
	Hex bool

	NoFile bool
	NoLine bool

	Limits  Limits
	Filters Filters

	// evidence bag
	Evidence struct {
		Case string
		File string
		Mode BagMode
		Sign string
		Url  string
		Auth string
		Ecs  bool
		Hec  bool
	}

	// ai flags
	AI struct {
		Query  string
		Model  string
		Embed  string
		NumCtx int
		Temp   float64
		TopP   float64
		TopK   int
		Seed   int
	}

	Deflate struct {
		Pass string
	}

	// optional flags
	Optional struct {
		Raw       bool
		Readonly  bool
		NoConvert bool
		NoDeflate bool
		NoPlugins bool
	}

	// alias flags
	Alias struct {
		Logstash bool
		Splunk   bool
		Text     bool
		Json     bool
		Jsonl    bool
		Sqlite   bool
	}

	// entropy command
	Entropy struct {
		Min float64
		Max float64
	}

	// hash command
	Hash struct {
		Algos Algorithms
	}

	// strings command
	Strings struct {
		Re    *regexp.Regexp
		Min   int
		Max   int
		Class bool
	}
}

var flg *Flags = nil // singleton

func Get() *Flags {
	if flg == nil {
		flg = new(Flags)

		// set defaults
		flg.Evidence.Mode = BagModeText
	}

	return flg
}
