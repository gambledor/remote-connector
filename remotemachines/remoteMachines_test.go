// Package file
package remotemachines

import (
	"fmt"
	"os"
	"testing"
)

func TestReadConfigFileWithSuccess(t *testing.T) {
	t.Log("Given the need to test file configuration reading.")

	var dir = os.Getenv("HOME")
	var configFile = ".remote_connections"

	t.Logf("When given %s/%s for reading file", dir, configFile)
	var got = RemoteMachines{}
	err := got.ReadConfigFile(dir, configFile)
	if err != nil {
		panic(err)
	}
	t.Log("Should be able to read the file.")
	if len(got.Machines) == 0 {
		t.Error("Got an empty configuration")
	}
	t.Log("Should get a non empty content.")
}

func TestReadConfigFileWithNoFileFound(t *testing.T) {
	t.Log("Given the need to test configuration file reading without any file to read")

	var path = os.Getenv("HOME")
	var fileName = "itdoesnotexist"

	t.Logf("When given %s/%s for reading file.", path, fileName)
	expected := fmt.Sprintf("open %s/%s: no such file or directory", path, fileName)
	var remoteMachines = RemoteMachines{}
	err := remoteMachines.ReadConfigFile(path, "itdoesnotexist")
	if err.Error() != expected {
		t.Error(err.Error())
		t.Error("The config file should have been not found")
	}
	t.Log("Should get no file.")
}
