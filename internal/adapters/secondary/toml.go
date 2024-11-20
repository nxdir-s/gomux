package secondary

import (
	"os"

	"github.com/nxdir-s/gomux/internal/core/entity"
	"github.com/nxdir-s/gomux/internal/core/entity/config"
	"github.com/pelletier/go-toml/v2"
)

const (
	ConfigFile string = ".gomux.toml"
)

type ErrReadCfg struct {
	err error
}

func (e *ErrReadCfg) Error() string {
	return "failed to read " + ConfigFile + ": " + e.err.Error()
}

type ErrUnmarshalToml struct {
	err error
}

func (e *ErrUnmarshalToml) Error() string {
	return "failed to unmarshal " + ConfigFile + ": " + e.err.Error()
}

type Config struct {
	Session    string `toml:"session"`
	Project    string `toml:"project"`
	StartIndex int    `toml:"start_index"`
	Windows    map[any]Window
}

type Window struct {
	Name string `toml:"name"`
	Cmd  string `toml:"cmd"`
}

type TomlAdapter struct {
	cfg *Config
}

func NewTomlAdapter() (*TomlAdapter, error) {
	return &TomlAdapter{}, nil
}

func (a *TomlAdapter) LoadConfig() (*entity.Config, error) {
	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		return nil, &ErrReadCfg{err}
	}

	err = toml.Unmarshal(data, a.cfg)
	if err != nil {
		return nil, &ErrUnmarshalToml{err}
	}

	windows := make([]config.Window, 0, len(a.cfg.Windows))
	for _, window := range a.cfg.Windows {
		windows = append(windows, config.Window{
			Name: window.Name,
			Cmd:  window.Cmd,
		})
	}

	return &entity.Config{
		Session:    a.cfg.Session,
		Project:    a.cfg.Project,
		StartIndex: a.cfg.StartIndex,
		Windows:    windows,
	}, nil
}
