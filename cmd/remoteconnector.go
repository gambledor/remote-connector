// Copyright (c) 2018-present, Giuseppe Lo Brutto All rights reserved

// Package main provides the command to make remote connection
// by a configuration hidden file name .remote_connections
// Before build exec git rev-parse --short HEAD and place
// the result in build command as follows
// go build -ldflags "-x main.Build=<sha1>" cmd/remoteconnector.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gambledor/remote-connector/remotemachine"
)

const (
	// ConfFileName is the configuration file name
	ConfFileName string = ".remote_connections"
	// version is the software version
	version string = "0.5"
	// author is the software author
	author = "Giuseppe Lo Brutto"
)

var (
	// Build is to compile passing -ldflags "-X main.Build <build sha1>"
	Build  string
	choise int
)

func init() {
	flag.IntVar(&choise, "c", 0, "The chosen remote machine")
}

func showRemoteMachines(remoteMachines *[]remotemachine.RemoteMachine) {
	for idx, machine := range *remoteMachines {
		// add 1 as index offset
		fmt.Printf("%d) %s\n", idx+1, machine.Name)
	}
	// press 0 to quit
	fmt.Printf("%d) %s\n", 0, "quit")
}

func getChoice(remoteMachines *[]remotemachine.RemoteMachine) int {
	var choise int
	var exit bool // initialized to false
	// 3. the user makes a choise to witch machine wants to connect to
	for !exit {
		fmt.Print("> ")
		var err error
		var input string
		if _, err = fmt.Scanf("%s", &input); err != nil {
			fmt.Println("No choise has been made")
		}
		if choise, err = strconv.Atoi(input); err != nil {
			fmt.Println("You have to enter a number")
		}
		if err == nil && choise >= 0 && choise <= len(*remoteMachines) {
			exit = true
		}
	}

	return choise
}

func main() {

	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("remote connector \033[32m%s-%s\033[0m, created by \033[96m%s\n", version, Build, author)
		os.Exit(0)
	}

	// 1. read configuration file for remote connections
	var remoteMachines *[]remotemachine.RemoteMachine

	remoteMachines, err := remotemachine.ReadConfigFile(os.Getenv("HOME"), ConfFileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if len(os.Args) > 1 && os.Args[1] == "list" {
		showRemoteMachines(remoteMachines)
		os.Exit(0)
	}
	flag.Parse()
	// 2. show a remote connections menu and get choise
	if choise == 0 {
		showRemoteMachines(remoteMachines)
		choise = getChoice(remoteMachines)
	}
	// 4. make a ssh connction to the chosen machine
	if choise > 0 && choise <= len(*remoteMachines) {
		fmt.Printf("You've chosen to connect to \033[96m%s\033[0m\n", (*remoteMachines)[choise-1].Host)
		if err := (*remoteMachines)[choise-1].Connect(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Bye Bye")
}
