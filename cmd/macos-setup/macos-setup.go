package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/vrunoa/macos-setup/internal/brew"
	"github.com/vrunoa/macos-setup/internal/config"
	"github.com/vrunoa/macos-setup/internal/exec"
	"github.com/vrunoa/macos-setup/internal/setup"
	"github.com/vrunoa/macos-setup/internal/ui"
	"github.com/vrunoa/macos-setup/internal/version"
)

func setupLogging(verbose bool) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"})
}

var configFile string
var commander exec.Cmd

func installCommand(commander exec.Cmd) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "install",
		Run: func(cmd *cobra.Command, args []string) {
			brew, err := brew.New(commander, configFile)
			fmt.Println(brew)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to setup brew")
			}
			brew.InstallFormulas()
		},
	}
	cmd.Flags().StringVarP(&configFile, "config-file", "c", "", "config-file")
	return cmd
}

func setupCommand(commander exec.Cmd) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "setup",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.New(configFile)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to read config file")
			}
			log.Info().Msg("Welcome to your setup helper")
			fmt.Print(ui.Floppy)
			fmt.Println()
			setupHelper := setup.New(cfg)
			setupHelper.Run()
		},
	}
	cmd.Flags().StringVarP(&configFile, "config-file", "c", "", "config-file")
	return cmd
}

func uninstallCommand(commander exec.Cmd) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall",
		Run: func(cmd *cobra.Command, args []string) {
			brew, err := brew.New(commander, configFile)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to setup brew")
			}
			brew.UninstallFormulas()
		},
	}
	cmd.Flags().StringVarP(&configFile, "config-file", "c", "", "config-file")
	return cmd
}

func main() {
	setupLogging(true)
	commander := exec.New()
	mainCmd := &cobra.Command{
		Use:     "macos-setup [command]",
		Short:   "CLI tool for setting up your macOS laptop",
		Version: fmt.Sprintf("%s\n(build %s)", version.Version, version.GitCommit),
	}
	mainCmd.AddCommand(setupCommand(commander))
	mainCmd.AddCommand(installCommand(commander))
	mainCmd.AddCommand(uninstallCommand(commander))
	if err := mainCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("wops! seems like we messed up")
		os.Exit(1)
	}
}
