package setup

import (
	"fmt"
	"io"
	"os"
	"path"

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
	return nil
}

func (s *setupHelper) CopyFiles() error {
	files := s.Config.Files
	if len(files) == 0 {
		log.Info().Msg("No user files found. Skipping copy")
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
