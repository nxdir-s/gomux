package secondary

import (
	"fmt"
	"os"

	"github.com/nxdir-s/gomux/internal/core/entity"
	"github.com/nxdir-s/gomux/internal/core/entity/config"
	"github.com/pelletier/go-toml/v2"
)

type ErrReadCfg struct {
	err error
}

func (e *ErrReadCfg) Error() string {
	return "failed to read " + config.FileName + ": " + e.err.Error()
}

type ErrUnmarshalToml struct {
	err error
}

func (e *ErrUnmarshalToml) Error() string {
	return "failed to unmarshal " + config.FileName + ": " + e.err.Error()
}

type Config struct {
	Session    string `toml:"session"`
	StartIndex int    `toml:"start_index"`

	Windows map[any]Window
}

type Window struct {
	Name string   `toml:"name"`
	Cmd  []string `toml:"cmd"`
}

type TomlAdapter struct{}

// NewTomlAdapter creates a toml adapter
func NewTomlAdapter() (*TomlAdapter, error) {
	return &TomlAdapter{}, nil
}

// LoadConfig attempts to read a toml file in the current directory and returns a config
func (a *TomlAdapter) LoadConfig() (*entity.Config, error) {
	data, err := os.ReadFile(config.FileName)
	if err != nil {
		return nil, &ErrReadCfg{err}
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, &ErrUnmarshalToml{err}
	}

	fmt.Fprintf(os.Stdout, "config: %+v\n", cfg)

	windows := make([]config.Window, 0, len(cfg.Windows))

	for i := range cfg.Windows {
		windows = append(windows, config.Window{
			Name: cfg.Windows[i].Name,
			Cmd:  cfg.Windows[i].Cmd,
		})
	}

	return &entity.Config{
		Session:    cfg.Session,
		StartIndex: cfg.StartIndex,
		Windows:    windows,
	}, nil
}
