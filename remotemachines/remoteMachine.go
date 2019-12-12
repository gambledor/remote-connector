package remotemachines

import (
	"fmt"
	"os"
	"os/exec"
)

// RemoteMachine is a remote machine config
type RemoteMachine struct {
	User     string `json:"user"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Protocol string `json:"protocol"`
}

// Connect execute the connection to the chosen remoteMachine
func (rm RemoteMachine) Connect(withX bool) error {
	var cmd *exec.Cmd

	connectionString := fmt.Sprintf("%s@%s", rm.User, rm.Host)
	if withX {
		cmd = exec.Command(rm.Protocol, "-X", connectionString)
	} else {
		cmd = exec.Command(rm.Protocol, connectionString)
	}

	cmd = exec.Command(rm.Protocol, connectionString)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
