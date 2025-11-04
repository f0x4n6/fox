package user

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/cuhsat/fox/v3/internal/pkg/sys/fs"
)

func LoadConfig(cfg *viper.Viper, name string) bool {
	var ok bool

	cfg.SetConfigPermissions(0600)
	cfg.SetConfigType("toml")
	cfg.SetConfigName(name)

	for _, path := range []string{
		// system config
		"/etc/fox",
		// local config
		"/usr/local/etc/fox",
		// user config
		"$HOME/.config/fox",
	} {
		cfg.AddConfigPath(path)

		if cfg.MergeInConfig() == nil {
			ok = true
		}
	}

	return ok
}

func SaveConfig(cfg *viper.Viper, name string) bool {
	path := Config(name)

	err := os.MkdirAll(filepath.Dir(path), 0700)

	if err != nil {
		return false
	}

	err = cfg.WriteConfigAs(path)

	if err != nil {
		return false
	}

	return true
}

func TempDir(prefix string) string {
	tmp, err := os.MkdirTemp(Cache(), fmt.Sprintf("%s-*", prefix))

	if err != nil {
		log.Panicln(err)
	}

	return tmp
}

func TempFile(prefix string) *os.File {
	tmp, err := os.CreateTemp(Cache(), fmt.Sprintf("%s-*", prefix))

	if err != nil {
		log.Panicln(err)
	}

	return tmp
}

func Persist(name string) string {
	f, ok := fs.Open(name).(fs.File)

	if !ok {
		return name // regular file
	}

	t := TempFile("fox")

	_, err := t.WriteTo(f)

	if err != nil {
		log.Println(err)
	}

	return f.Name()
}

func Config(name string) string {
	dir, err := os.UserHomeDir()

	if err != nil {
		log.Println(err)
	}

	return filepath.Join(dir, ".config", "fox", name)
}

func Cache() string {
	dir, err := os.UserHomeDir()

	if err != nil {
		log.Println(err)
	}

	tmp := filepath.Join(dir, ".cache", "fox")

	err = os.MkdirAll(tmp, 0700)

	if err != nil {
		log.Println(err)
	}

	return tmp
}
