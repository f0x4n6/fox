package user

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/cuhsat/fox/v4/internal/pkg/sys/fs"
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
