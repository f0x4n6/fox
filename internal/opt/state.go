package opt

import (
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/gdamore/tcell/v2"

	"github.com/cuhsat/fox/v3/internal/pkg/flags"
	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
	"github.com/cuhsat/fox/v3/internal/pkg/text"
	"github.com/cuhsat/fox/v3/internal/pkg/types/mode"
	"github.com/cuhsat/fox/v3/internal/pkg/user/config"
)

type Listable func(string) []fs.Item
type Findable func(string) string

type State struct {
	sync.RWMutex

	Root tcell.Screen
	Icon *text.Icon

	List Listable
	Find Findable

	mode mode.Mode
	last mode.Mode

	path string

	model string
	embed string
	theme string

	space atomic.Uint32

	n atomic.Bool
	w atomic.Bool
	y atomic.Bool
	r atomic.Bool
}

func NewState(root tcell.Screen) *State {
	cfg := config.Get()
	flg := flags.Get()

	state := &State{
		// screen
		Root: root,

		// icons
		Icon: text.Icons(!flg.UI.Legacy),

		// calls
		List: fs.List,

		// modes
		mode: mode.Default,
		last: mode.Default,

		// models
		model: cfg.GetString("ai.model"),
		embed: cfg.GetString("ai.embed"),

		// theme
		theme: cfg.GetString("ui.theme"),
	}

	state.path, _ = os.Getwd()

	state.space.Store(cfg.GetUint32("ui.space"))

	state.n.Store(cfg.GetBool("ui.state.n"))
	state.w.Store(cfg.GetBool("ui.state.w"))
	state.y.Store(cfg.GetBool("ui.state.y"))
	state.r.Store(cfg.GetBool("ui.state.r"))

	// precede flags
	s := strings.ToUpper(flg.UI.State)

	if strings.ContainsRune(s, '-') {
		state.n.Store(false)
		state.w.Store(false)
		state.y.Store(false)
		state.r.Store(false)
	} else if len(s) > 0 {
		state.n.Store(strings.ContainsRune(s, 'N'))
		state.w.Store(strings.ContainsRune(s, 'W'))
		state.y.Store(strings.ContainsRune(s, 'Y'))
		state.r.Store(strings.ContainsRune(s, 'R'))
	}

	return state
}

func (s *State) Mode() mode.Mode {
	s.RLock()
	defer s.RUnlock()
	return s.mode
}

func (s *State) Last() mode.Mode {
	s.RLock()
	defer s.RUnlock()
	return s.last
}

func (s *State) Path() string {
	s.RLock()
	defer s.RUnlock()
	return s.path
}

func (s *State) Model() string {
	s.RLock()
	defer s.RUnlock()
	return s.model
}

func (s *State) Embed() string {
	s.RLock()
	defer s.RUnlock()
	return s.embed
}

func (s *State) Theme() string {
	s.RLock()
	defer s.RUnlock()
	return s.theme
}

func (s *State) Space() int {
	return int(s.space.Load())
}

func (s *State) IsNavi() bool {
	return s.n.Load()
}

func (s *State) IsWrap() bool {
	return s.w.Load()
}

func (s *State) IsSticky() bool {
	if s.Mode() == mode.Less {
		return s.y.Load()
	} else {
		return false
	}
}

func (s *State) IsReload() bool {
	return s.r.Load()
}

func (s *State) ForceRender() {
	_ = s.Root.PostEvent(tcell.NewEventInterrupt(nil))
}

func (s *State) SwitchMode(m mode.Mode) bool {
	// deny goto in static modes
	if m == mode.Goto && s.Mode().IsStatic() {
		return false
	}

	// react only to mode changes
	if m == s.Mode() {
		return false
	}

	s.Lock()
	s.last = s.mode
	s.mode = m
	s.Unlock()

	return true
}

func (s *State) ChangePath(p string) {
	s.Lock()
	s.path = p
	s.Unlock()
}

func (s *State) ChangeModel(m string) {
	s.Lock()
	s.model = m
	s.Unlock()
}

func (s *State) ChangeEmbed(e string) {
	s.Lock()
	s.embed = e
	s.Unlock()
}

func (s *State) ChangeTheme(t string) {
	s.Lock()
	s.theme = t
	s.Unlock()
}

func (s *State) SetListable(fn Listable) {
	s.Lock()
	s.List = fn
	s.Unlock()
}

func (s *State) SetFindable(fn Findable) {
	s.Lock()
	s.Find = fn
	s.Unlock()
}

func (s *State) ToggleNavi() {
	s.n.Store(!s.n.Load())
}

func (s *State) ToggleWrap() {
	s.w.Store(!s.w.Load())
}

func (s *State) ToggleSticky() {
	s.y.Store(!s.y.Load())
}

func (s *State) ToggleReload() {
	s.r.Store(!s.r.Load())
}

func (s *State) Call(fn func()) {
	go func() {
		fn()
		s.ForceRender()
	}()
}

func (s *State) Save() {
	cfg := config.Get()

	cfg.Set("ai.model", s.Model())
	cfg.Set("ai.embed", s.Embed())
	cfg.Set("ui.theme", s.Theme())
	cfg.Set("ui.space", s.Space())
	cfg.Set("ui.state.n", s.IsNavi())
	cfg.Set("ui.state.w", s.IsWrap())
	cfg.Set("ui.state.y", s.IsSticky())
	cfg.Set("ui.state.r", s.IsReload())

	if !flags.Get().Optional.Readonly {
		config.Save()
	}
}
