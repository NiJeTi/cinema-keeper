package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func ReadConfig[T any]() (*T, error) {
	k := koanf.New(".")

	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Println("config file has not been loaded")
	}

	cb := func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), "__", ".")
	}
	if err := k.Load(env.Provider("", ".", cb), nil); err != nil {
		log.Println("environment variables have not been loaded")
	}

	cfg := new(T)
	err := k.UnmarshalWithConf("", cfg, koanf.UnmarshalConf{Tag: "conf"})
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return cfg, nil
}
