package brew

import (
	"github.com/rs/zerolog/log"
	"github.com/vrunoa/macos-setup/internal/config"
	"github.com/vrunoa/macos-setup/internal/exec"
	"github.com/vrunoa/macos-setup/internal/yaml"
)

type brew struct {
	Config    config.Config
	Commander exec.Cmd
}

func (b *brew) getFormulas() []string {
	return b.Config.Brew.Formulas
}

func (b *brew) UninstallFormulas() {
	log.Info().Msg("uninstalling formulas")
	formulas := b.getFormulas()
	for _, formula := range formulas {
		log.Debug().Msgf("uninstalling -> %s", formula)
		err := b.Commander.Run("arch", []string{"-arm64", "brew", "uninstall", formula, "-d"})
		if err != nil {
			log.Warn().Err(err).Msgf("failed to install formulat -> %s", formula)
		}
	}
}

func (b *brew) InstallFormulas() {
	log.Info().Msg("installing formulas")
	formulas := b.getFormulas()
	for _, formula := range formulas {
		log.Debug().Msgf("installing -> %s", formula)
		err := b.Commander.Run("arch", []string{"-arm64", "brew", "install", formula, "-d"})
		if err != nil {
			log.Warn().Err(err).Msgf("failed to install formulat -> %s", formula)
		}
	}
}

func checkBrewInstalled(commander exec.Cmd) error {
	return commander.Run("brew", []string{"--version"})
}

func New(commander exec.Cmd, configFile string) (*brew, error) {
	err := checkBrewInstalled(commander)
	if err != nil {
		return nil, err
	}
	var config config.Config
	err = yaml.ReadYaml(&config, configFile)
	if err != nil {
		return nil, err
	}
	return &brew{
		Config:    config,
		Commander: commander,
	}, nil
}
