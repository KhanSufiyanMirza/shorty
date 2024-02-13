package config

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var ConfigFile = flag.String("configPath", "./config.yaml", "The environment configuaration path")

type Provider struct {
	Config    *koanf.Koanf
	YamlFiles []string
	JsonFiles []string
}

type Config struct {
	AppConfig
}

func NewAppConfig() Config {
	flag.Parse()
	if *ConfigFile == "" {
		log.Fatal("unable to load environment config")
	}

	var configFiles = []string{*ConfigFile}
	var provider = getProvider(configFiles...)
	return initAppConfig(provider)
}

func getProvider(files ...string) *Provider {
	return &Provider{
		Config:    koanf.New("."),
		YamlFiles: files,
		JsonFiles: nil,
	}
}

func initAppConfig(ko *Provider) Config {

	ko.Yaml()
	ko.Env()
	var appConfig Config
	err := ko.Config.UnmarshalWithConf("", &appConfig, koanf.UnmarshalConf{Tag: "json"})
	if err != nil {
		log.Fatal("error while unmarshalling config")
	}

	return appConfig
}

func (c *Provider) Yaml() {
	for _, f := range c.YamlFiles {
		if err := c.Config.Load(file.Provider(f), yaml.Parser()); err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("config file not found. Please pass a configuration file. %v", err)
			}
			log.Fatalf("error loading config from file: %v.", err)
		}
	}
}

func (c *Provider) Env() {
	if err := c.Config.Load(env.Provider("ENV_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "ENV_")), "_", ".", -1)
	}), nil); err != nil {
		log.Fatalf("error loading config from env: %v", err)
	}
}

func (c Config) Validate(isRDB, isKeyDB, isEventProducer, isEventConsumer bool) {
	var err error
	if isRDB {
		if err = c.AppConfig.ValidateRDBConfig(); err != nil {
			log.Fatalf("error in RDB config, %v", err)
		}
	}
	if isKeyDB {
		if err = c.AppConfig.ValidateKeyDBConfig(); err != nil {
			log.Fatalf("error in KeyDB config, %v", err)
		}
	}

	if isEventProducer {
		if err = c.AppConfig.ValidateEventProducerConfig(); err != nil {
			log.Fatalf("error in EventProducer config, %v", err)
		}
	}
	if isEventConsumer {
		if err = c.AppConfig.ValidateEventConsumerConfig(); err != nil {
			log.Fatalf("error in EventConsumer config, %v", err)
		}
	}
}
