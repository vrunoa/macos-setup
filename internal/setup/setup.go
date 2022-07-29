package setup

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rs/zerolog/log"
	"github.com/vrunoa/macos-setup/internal/config"
)

type setupHelper struct {
	Config config.Config
}

func (s *setupHelper) Run() error {
	err := s.CopyFiles()
	if err != nil {
		return err
	}
	s.InstallInteractive()
	return nil
}

func (s *setupHelper) InstallInteractive() {
	apps := s.Config.Applications.Interactive
	if len(apps) == 0 {
		log.Info().Msg("No interactive install defined. Skipping")
		return
	}
	log.Info().Msg("Found interactive installs. Open a new terminal to run some commands")
	for _, app := range apps {
		fmt.Println()
		fmt.Println(fmt.Sprintf("\tApplication: %s", app.Name))
		fmt.Println(fmt.Sprintf("\tHome: %s", app.Home))
		fmt.Println("\tCommand:")
		fmt.Println(fmt.Sprintf("\t%s", app.Cmd))
		fmt.Println()
		keepGoing := false
		for {
			time.Sleep(3000)
			promt := &survey.Confirm{
				Message: "Is installation done. Want to continue to?",
			}
			survey.AskOne(promt, &keepGoing)
			if keepGoing {
				break
			}
		}
	}
}

func (s *setupHelper) CopyFiles() error {
	files := s.Config.Files
	if len(files) == 0 {
		log.Info().Msg("No user files defined. Skipping copy")
		return nil
	}
	log.Info().Msg("Found user files")
	for _, f := range files {
		fmt.Println(fmt.Sprintf("\t * %s", f))
	}
	shouldCopy := false
	prompt := &survey.Confirm{
		Message: "Do you want to copy this file to your home folder?",
	}
	survey.AskOne(prompt, &shouldCopy)
	if !shouldCopy {
		return nil
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	for _, f := range files {
		fPath := path.Join(curDir, "files", f)
		log.Info().Msgf("copying file -> %s", fPath)
		src, err := os.Open(fPath)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to read file -> %s", fPath)
		}
		defer src.Close()
		dstPath := path.Join(homeDir, f)
		dst, err := os.Create(dstPath)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to create file -> %s", dstPath)
		}
		defer dst.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to copy file -> %s", fPath)
		}
	}
	return nil
}

func New(cfg *config.Config) *setupHelper {
	return &setupHelper{
		Config: *cfg,
	}
}
