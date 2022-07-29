package exec

import (
	"bufio"
	"fmt"
	sysExec "os/exec"
)

type Cmd interface {
	Run(cmd string, args []string) error
}

type cmd struct{}

func (c *cmd) Run(cmd string, args []string) error {
	command := sysExec.Command(cmd, args...)
	stderr, _ := command.StderrPipe()
	err := command.Start()
	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	command.Wait()
	return err
}

func New() *cmd {
	return &cmd{}
}
