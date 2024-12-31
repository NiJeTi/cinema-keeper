package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func ReadConfig[T any]() (*T, error) {
	k := koanf.New(".")

	//nolint:revive,staticcheck // some action might be needed
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		// config file has not been loaded
	}

	cb := func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), "__", ".")
	}
	//nolint:revive,staticcheck // some action might be needed
	if err := k.Load(env.Provider("", ".", cb), nil); err != nil {
		// environment variables have not been loaded
	}

	cfg := new(T)
	err := k.UnmarshalWithConf("", cfg, koanf.UnmarshalConf{Tag: "conf"})
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return cfg, nil
}
