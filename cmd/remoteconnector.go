// Copyright (c) 2018-present, Giuseppe Lo Brutto All rights reserved

// Package main provides the command to make remote connection
// by a configuration hidden file name .remote_connections
// Before build exec git rev-parse --short HEAD and place
// the result in build command as follows
// go build -ldflags "-X main.Build=`git rev-parse --short HEAD" cmd/remoteconnector.go
package main

import (
	"flag"
	"fmt"
	"github.com/gambledor/remote-connector/build"
	"github.com/gambledor/remote-connector/remotemachines"
	"log"
	"os"
)

const (
	// ConfFileName holds the configuration file name
	confFileName string = ".remote_connections"
	version             = "Version: remote connector \033[32m%s-%s\033[0m, created by \033[96m%s\033[0m\nbuilt by \033[96m%s\033[0m\non %s\n"
	usage        string = `Usage: %s
Run remoteconnector

Options:
`
)

var (
	choice int
	xmode  bool
)

func init() {
	flag.IntVar(&choice, "c", 0, "The chosen remote machine.")
	flag.BoolVar(&xmode, "X", false, "Enable X mode.")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}
}

func manageConnection(remoteMachines remotemachines.RemoteMachines) {
	// 2. show a remote connections menu and get choice
	if choice == 0 {
		remoteMachines.ShowRemoteMachinesMenu()
		choice = remoteMachines.GetChoice()
	}
	// 3. make a ssh connction to the chosen machine
	if choice > 0 && choice <= len(remoteMachines.Machines) {
		fmt.Printf("You've chosen to connect to \033[96m%s\033[0m\n", remoteMachines.Machines[choice-1].Host)
		if err := remoteMachines.Machines[choice-1].Connect(xmode); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Bye Bye")
}

func main() {
	flag.Parse()

	// 1. read configuration file for remote connections
	var remoteMachines remotemachines.RemoteMachines
	if err := remoteMachines.ReadConfigFile(os.Getenv("HOME"), confFileName); err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 1 {
		manageConnection(remoteMachines)
	} else {
		switch os.Args[1] {
		case "version":
			fmt.Printf(version, build.Version, build.Build, build.Author, build.User, build.Time)
			os.Exit(0)
		case "list":
			remoteMachines.ShowRemoteMachinesMenu()
			os.Exit(0)
		default:
			log.Fatalf("Error: unknown command %s", os.Args[1])
		}
	}
}
