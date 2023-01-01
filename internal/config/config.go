package config

import "github.com/vrunoa/macos-setup/internal/yaml"

type Application struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd,omitempty"`
	Home string `yaml:"home"`
}

type Applications struct {
	Interactive []Application `yaml:"interactive"`
	Manual      []Application `yaml:"manual"`
}

type Brew struct {
	Formulas []string `yaml:"formulas"`
}

type NVM struct {
	Node []string `yaml:"nodeVersions"`
}

type NPM struct {
	Packages []string `yaml:"packages"`
}

type Config struct {
	Files        []string     `yaml:"files"`
	Applications Applications `yaml:"applications"`
	Brew         Brew         `yaml:"brew"`
	NVM          NVM          `yaml:"nvm"`
	NPM          NPM          `yaml:"npm"`
}

func New(configFile string) (*Config, error) {
	var cfg Config
	err := yaml.ReadYaml(&cfg, configFile)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
