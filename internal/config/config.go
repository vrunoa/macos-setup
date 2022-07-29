package config

import "github.com/vrunoa/macos-setup/internal/yaml"

type application struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd,omitempty"`
	Home string `yaml:"home"`
}

type applications struct {
	Interactive []application `yaml:"interactive"`
	Manual      []application `yaml:"manual"`
}

type brew struct {
	Formulas []string `yaml:"formulas"`
}

type nvm struct {
	Node []string `yaml:"nodeVersions"`
}

type npm struct {
	Packages []string `yaml:"packages"`
}

type Config struct {
	Files        []string     `yaml:"files"`
	Applications applications `yaml:"applications"`
	Brew         brew         `yaml:"brew"`
	NVM          nvm          `yaml:"nvm"`
	NPM          npm          `yaml:"npm"`
}

func New(configFile string) (*Config, error) {
	var cfg Config
	err := yaml.ReadYaml(&cfg, configFile)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
