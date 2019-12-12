// Package remotemachine provides file management utils
// Copyright (c) 2018-present, Giuseppe Lo Brutto All rights reserved
package remotemachines

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// RemoteMachines is a list of remote machine connections
type RemoteMachines struct {
	Machines []RemoteMachine `json:"machines"`
}

// ReadConfigFile reads the remote machine file config
func (rm *RemoteMachines) ReadConfigFile(basePath, fileName string) error {
	file, err := ioutil.ReadFile(filepath.Join(basePath, fileName))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, &rm); err != nil {
		return err
	}

	return nil
}

func add(x, y int) int {
	return x + y
}

// ShowRemoteMachinesMenu shows the menu of configured remote connections.
func (rm RemoteMachines) ShowRemoteMachinesMenu() {
	const templateMenu = `------------------------------------------------------------
REMOTE MACHINES
------------------------------------------------------------
{{ range $index, $item := .Machines }} {{ add $index 1 }} - {{ $item.Name }}
{{ else }} no remote machines configured {{ end }}------------------------------------------------------------
Press 0 to quit.
------------------------------------------------------------
`
	var menu = template.Must(template.New("menu").Funcs(template.FuncMap{"add": add}).Parse(templateMenu))
	if err := menu.Execute(os.Stdout, rm); err != nil {
		log.Fatal(err)
	}
}

// GetChoice get the number of the remote machine chosen
func (rm RemoteMachines) GetChoice() int {
	var choice int
	var exit bool // initialized to false
	var err error
	var input string

	// 3. the user makes a choice to witch machine wants to connect to
	for !exit {
		fmt.Print("> ")
		if _, err = fmt.Scanf("%s", &input); err != nil {
			fmt.Println("No choice has been made")
		}
		if choice, err = strconv.Atoi(input); err != nil {
			fmt.Println("You have to enter a number")
		}
		if err == nil && choice >= 0 && choice <= len(rm.Machines) {
			exit = true
		}
	}

	return choice
}
