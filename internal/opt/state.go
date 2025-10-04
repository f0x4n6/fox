package opt

import (
	"strings"
	"sync"
	"sync/atomic"

	"github.com/gdamore/tcell/v2"

	"github.com/cuhsat/fox/internal/pkg/flags"
	"github.com/cuhsat/fox/internal/pkg/text"
	"github.com/cuhsat/fox/internal/pkg/types/mode"
	"github.com/cuhsat/fox/internal/pkg/user/config"
)

type State struct {
	sync.RWMutex

	Root tcell.Screen
	Icon *text.Icon

	mode mode.Mode
	last mode.Mode

	model string
	embed string
	theme string

	space atomic.Uint32

	n atomic.Bool
	w atomic.Bool
	t atomic.Bool
	p atomic.Bool
}

func NewState(root tcell.Screen) *State {
	cfg := config.Get()
	flg := flags.Get()

	state := &State{
		// screen
		Root: root,

		// icons
		Icon: text.Icons(!flg.UI.Legacy),

		// modes
		mode: mode.Default,
		last: mode.Default,

		// models
		model: cfg.GetString("ai.model"),
		embed: cfg.GetString("ai.embed"),

		// theme
		theme: cfg.GetString("ui.theme"),
	}

	state.space.Store(cfg.GetUint32("ui.space"))

	state.n.Store(cfg.GetBool("ui.state.n"))
	state.w.Store(cfg.GetBool("ui.state.w"))
	state.t.Store(cfg.GetBool("ui.state.t"))
	state.p.Store(false)

	// precede flags
	s := strings.ToUpper(flg.UI.State)

	if strings.ContainsRune(s, '-') {
		state.n.Store(false)
		state.w.Store(false)
		state.t.Store(false)
	} else if len(s) > 0 {
		state.n.Store(strings.ContainsRune(s, 'N'))
		state.w.Store(strings.ContainsRune(s, 'W'))
		state.t.Store(strings.ContainsRune(s, 'T'))
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

func (s *State) IsFollow() bool {
	return s.t.Load()
}

func (s *State) IsPinned() bool {
	if s.Mode() == mode.Less {
		return s.p.Load()
	} else {
		return false
	}
}

func (s *State) ForceRender() {
	_ = s.Root.PostEvent(tcell.NewEventInterrupt(nil))
}

func (s *State) SwitchMode(m mode.Mode) bool {
	// deny goto in static modes
	if m == mode.Goto && s.Mode().Static() {
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

func (s *State) ToggleNavi() {
	s.n.Store(!s.n.Load())
}

func (s *State) ToggleWrap() {
	s.w.Store(!s.w.Load())
}

func (s *State) ToggleFollow() {
	s.t.Store(!s.t.Load())
}

func (s *State) TogglePinned() {
	s.p.Store(!s.p.Load())
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
	cfg.Set("ui.state.t", s.IsFollow())

	if !flags.Get().Opt.Readonly {
		config.Save()
	}
}
