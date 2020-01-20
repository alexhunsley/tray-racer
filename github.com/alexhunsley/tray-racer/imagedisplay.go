package main

import (
	"log"
	"os"
	"os/exec"
)

// show a specified file by Preview.app for OS X(darwin)
func Show(name string) {
	previewCommandPath := "/System/Applications/Preview.app/Contents/MacOS/Preview"

	if _, err := os.Stat(previewCommandPath); os.IsNotExist(err) {
		previewCommandPath = "/Applications/Preview.app/Contents/MacOS/Preview"
	}
	command := "open"
	arg1 := "-a"
	//arg2 :=
	cmd := exec.Command(command, arg1, previewCommandPath, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
