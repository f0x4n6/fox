package flags

var CLI struct {
	// File infos
	Hex     bool    `short:"x"`
	Count   bool    `short:"w"`
	Hash    string  `short:"a" sep:"," default:"SHA256"`
	Entropy float32 `short:"y" sep:":"`
	Strings int     `short:"s" sep:":" default:"3"`

	// File limits
	Head  bool `short:"h" xor:"head,tail"`
	Tail  bool `short:"t" xor:"head,tail"`
	Lines int  `short:"n" default:"10"`
	Bytes int  `short:"c" default:"16"`

	// File loader
	Pass string `short:"p"`

	// Line filter
	Regex   string `short:"e"`
	Context int    `short:"C"`
	Before  int    `short:"B"`
	After   int    `short:"A"`

	// LLM parser
	Query string `short:"q" default:"analyse"`
	Model string `short:"m"`
	Embed string

	// LLM options
	NumCtx int     `default:"4096"`
	Temp   float32 `default:"0.2"`
	TopP   float32 `default:"0.5"`
	TopK   int     `default:"10"`
	Seed   int     `default:"8211"`

	// Evidence bag
	Case string `short:"N"`
	File string `short:"F" default:"evidence"`
	Mode string `default:"text"`

	// Evidence sign
	Sign string

	// Evidence URL
	Url  string `short:"u"`
	Auth string
	Ecs  bool `xor:"ecs,hec"`
	Hec  bool `xor:"ecs,hec" and:"hec,auth"`

	// Turn off
	Readonly  bool `short:"R"`
	Raw       bool `short:"r"`
	NoFile    bool `long:"no-file"`
	NoLine    bool `long:"no-line"`
	NoConvert bool `long:"no-convert"`
	NoDeflate bool `long:"no-deflate"`

	// Alias
	Logstash bool `short:"L"`
	Splunk   bool `short:"S"`
	Text     bool `short:"T"`
	Json     bool `short:"j"`
	Jsonl    bool `short:"J"`
	Sqlite   bool `short:"Q"`

	// Standard
	Help    bool
	Version bool

	// Positional arguments
	Paths []string `arg:"" name:"path" type:"path" optional:""`
}
