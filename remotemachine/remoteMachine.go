// Package remotemachine provides file management utils
// Copyright (c) 2018-present, Giuseppe Lo Brutto All rights reserved
package remotemachine

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// RemoteMachine is a remote machine config
type RemoteMachine struct {
	User     string `json:user`
	Name     string `json:name`
	Host     string `json:host`
	Protocol string `json:protocol`
}

// ReadConfigFile reads the remote machine file config
func ReadConfigFile(basePath, fileName string) (*[]RemoteMachine, error) {
	file, err := ioutil.ReadFile(filepath.Join(basePath, fileName))
	if err != nil {
		return nil, err
	}

	var remoteMachines *[]RemoteMachine
	if err := json.Unmarshal(file, &remoteMachines); err != nil {
		return nil, err
	}

	return remoteMachines, nil
}

// Connect execute the connection to the chosen remoteMachine
func (rm *RemoteMachine) Connect(withX bool) error {
	connectionString := fmt.Sprintf("%s@%s", rm.User, rm.Host)

	var cmd *exec.Cmd
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

func add(x, y int) int {
	return x + y
}

// ShowRemoteMachinesMenu shows the menu of configured remote connections.
func ShowRemoteMachinesMenu(remoteMachines *[]RemoteMachine) {
	const templateMenu = `------------------------------------------------------------
REMOTE MACHINES
------------------------------------------------------------
{{ range $index, $item := . }} {{ add $index 1 }} - {{ $item.Name }}
{{ else }} no remote machines configured {{ end }}------------------------------------------------------------
Press 0 to quit.
------------------------------------------------------------
`
	var menu = template.Must(template.New("menu").Funcs(template.FuncMap{"add": add}).Parse(templateMenu))
	if err := menu.Execute(os.Stdout, remoteMachines); err != nil {
		log.Fatal(err)
	}
}
