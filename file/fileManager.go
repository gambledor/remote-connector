// Copyright (c) 2018-present, Giuseppe Lo Brutto All rights reserved

// Package file provide file management utils
package file

import (
	"encoding/json"
	"io/ioutil"
)

// RemoteMachine is a remote machine config
type RemoteMachine struct {
	User     string `json:user`
	Name     string `json:name`
	Host     string `json:host`
	Protocol string `json:protocol`
}

// ReadConfigFile reads the remote machine file config
func ReadConfigFile(filePath, fileName string) ([]RemoteMachine, error) {
	file, err := ioutil.ReadFile(filePath + "/" + fileName)
	if err != nil {
		return nil, err
	}

	var remoteMachines []RemoteMachine
	if err := json.Unmarshal(file, &remoteMachines); err != nil {
		return nil, err
	}

	return remoteMachines, nil
}
