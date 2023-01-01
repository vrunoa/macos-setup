package brew

import (
	"reflect"
	"testing"

	"github.com/vrunoa/macos-setup/internal/config"
)

type fakeCommander struct {
	RunFn func(cmd string, args []string) error
}

func (f *fakeCommander) Run(cmd string, args []string) error {
	return f.RunFn(cmd, args)
}

func TestInstall(t *testing.T) {
	var execCmd string
	var execArgs []string
	var execErr error
	commander := &fakeCommander{
		RunFn: func(cmd string, args []string) error {
			execCmd = cmd
			execArgs = args
			execErr = nil
			return execErr
		},
	}
	b := &brew{
		Config: config.Config{
			Brew: config.Brew{
				Formulas: []string{"some-package"},
			},
		},
		Commander: commander,
	}
	b.InstallFormulas()
	expectedCmd := "arch"
	expectedArgs := []string{"-arm64", "brew", "install", "some-package", "-d"}
	if execErr != nil {
		t.Errorf("error raised -> %v", execErr)
	}
	if execCmd != expectedCmd {
		t.Errorf("wrong cmd. Want: %v. Got: %v", expectedCmd, execCmd)
	}
	if !reflect.DeepEqual(execArgs, expectedArgs) {
		t.Errorf("wrong args. Want: %v. Got: %v", expectedArgs, execArgs)
	}
}

func TestUninstall(t *testing.T) {
	var execCmd string
	var execArgs []string
	var execErr error
	commander := &fakeCommander{
		RunFn: func(cmd string, args []string) error {
			execCmd = cmd
			execArgs = args
			execErr = nil
			return execErr
		},
	}
	b := &brew{
		Config: config.Config{
			Brew: config.Brew{
				Formulas: []string{"some-package"},
			},
		},
		Commander: commander,
	}
	b.UninstallFormulas()
	expectedCmd := "arch"
	expectedArgs := []string{"-arm64", "brew", "uninstall", "some-package", "-d"}
	if execErr != nil {
		t.Errorf("error raised -> %v", execErr)
	}
	if execCmd != expectedCmd {
		t.Errorf("wrong cmd. Want: %v. Got: %v", expectedCmd, execCmd)
	}
	if !reflect.DeepEqual(execArgs, expectedArgs) {
		t.Errorf("wrong args. Want: %v. Got: %v", expectedArgs, execArgs)
	}
}
