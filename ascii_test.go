package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	ascii "ascii/utilities"
)

func TestExecCommand(t *testing.T) {
	file, _ := os.ReadFile("standard.txt")
	fileString := ""
	fileData := strings.Split(string(file), "\n")

	// Create files
	f, _ := os.Create("testing.txt")
	g, _ := os.Create("output.txt")

	// close the files
	defer f.Close()
	defer g.Close()
	//
	testStrings := [1]string{
		"HeLlO\nthere 2+2=4",
	}
	// Define the commands
	commands := [1]*exec.Cmd{
		exec.Command("go", "run", ".", "HeLlO\nthere 2+2=4"),
	}

	//
	for i := range testStrings {
		inputParts, _ := ascii.HandleNewLine(testStrings[i])
		for _, part := range inputParts {
			if part == "" {
				fileString += "\n"
				continue
			}
			for i := 0; i < 8; i++ {
				for _, char := range part {
					startingIndex := ascii.GetStartingIndex(int(char))
					fileString += fileData[startingIndex-1+i]
				}
				fileString += "\n"
			}

		}
		fileString += "\n"
	}
	f.WriteString(fileString)

	// Execute each command and store it in a string
	for _, cmd := range commands {
		// commandline output
		output, err := cmd.Output()
		if err != nil {
			t.Error(err)
		}
		g.WriteString(string(output) + "\n")
		//
	}

	//
	test1, err := os.ReadFile("output.txt")
	if err != nil {
		t.Error(err)
	}
	test2, err := os.ReadFile("testing.txt")
	if err != nil {
		t.Error(err)
	}

	if string(test1) != string(test2) {
		t.Errorf("Output does not match. Expected %v got %v", string(test1), (string(test2)))
	}
}
