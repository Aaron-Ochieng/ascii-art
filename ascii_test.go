package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	ascii "ascii/utilities"
)

func TestExecCommand(t *testing.T) {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		t.Error(err)
		return
	}
	fileString := ""
	fileData := strings.Split(string(file), "\n")

	// Create files
	f, err := os.Create("testing.txt")
	if err != nil {
		t.Errorf("Error :%v", err)
	}
	g, err := os.Create("output.txt")
	if err != nil {
		t.Errorf("Error :%v", err)
	}
	// close the files
	defer f.Close()
	defer g.Close()

	// Modify here : Input your test input string
	testStrings := [1]string{
		"hello",
	}

	// Define the commands to execute here:
	commands := [1]*exec.Cmd{
		exec.Command("go", "run", ".", "hello"),
	}

	//
	for i := range testStrings {
		inputParts, err := ascii.HandleNewLine(testStrings[i])
		if err != nil {
			t.Error(err)
		}
		legalChar, err := ascii.CheckIllegalChar(inputParts) // Check for illegal characters in the string
		if err != nil {
			t.Error(err)
		}
		for _, part := range legalChar {
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

	// Read the output (1) of the commands from a file
	test1, err := os.ReadFile("output.txt")
	if err != nil {
		t.Error(err)
	}
	// Read the output (2) of the commands from a file
	test2, err := os.ReadFile("testing.txt")
	if err != nil {
		t.Error(err)
	}

	// Match the strings of the file
	if string(test1) != string(test2) {
		t.Errorf("Output does not match. Expected %v\n got %v\n", string(test1), (string(test2)))
	}
}
