package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func ReadConfig[T any]() T {
	k := koanf.New(".")

	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Println("config file has not been loaded")
	}

	cb := func(s string) string {
		return strings.Replace(strings.ToLower(s), "__", ".", -1)
	}
	if err := k.Load(env.Provider("", ".", cb), nil); err != nil {
		log.Println("environment variables have not been loaded")
	}

	var cfg T
	err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "conf"})
	if err != nil {
		log.Fatalln("failed to unmarshal config:", err)
	}
	return cfg
}
