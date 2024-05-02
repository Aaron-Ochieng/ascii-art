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
		os.Exit(1)
	}
	fileString := ""
	fileData := strings.Split(string(file), "\n")

	// Create files
	expected, err := os.Create("expected.txt")
	if err != nil {
		t.Errorf("Error :%v", err)
	}
	got, err := os.Create("got.txt")
	if err != nil {
		t.Errorf("Error :%v", err)
	}
	// close the files
	defer expected.Close()
	defer got.Close()

	// Modify here : Input your test input string
	testStrings := []string{
		"hello",
	}

	// Define the commands to execute here:
	commands := []*exec.Cmd{
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
					fileString += fileData[startingIndex+i]
				}
				fileString += "\n"
			}

		}
		fileString += "\n"
	}
	expected.WriteString(fileString)

	// Execute each command and store it in a string
	for _, cmd := range commands {
		// commandline output
		output, err := cmd.Output()
		if err != nil {
			t.Error(err)
		}
		got.WriteString(string(output) + "\n")
		//
	}

	// Read the output (1) of the commands from a file
	exTest, err := os.ReadFile("got.txt")
	if err != nil {
		t.Error(err)
	}
	// Read the output (2) of the commands from a file
	gotTest, err := os.ReadFile("expected.txt")
	if err != nil {
		t.Error(err)
	}

	// Match the strings of the file
	if string(exTest) != string(gotTest) {
		t.Errorf("Output does not match. Expected \n%v\n with length %v got \n%v\n with length %v", string(exTest), len(string(exTest)), (string(gotTest)), len(string(gotTest)))
	}
}
