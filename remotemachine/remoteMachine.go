// Copyright (c) 2018-present, Giuseppe Lo Brutto All rights reserved

// Package file provide file management utils
package remotemachine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
func (rm *RemoteMachine) Connect() error {
	connectionString := fmt.Sprintf("%s@%s", rm.User, rm.Host)
	cmd := exec.Command(rm.Protocol, connectionString)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
