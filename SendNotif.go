package main

import (
	"fmt"
	"os"
	"os/exec"
)

func SendNotif(message string) {
	command := "notify-send"

	args := []string{"-u", "critical", message} // Separate each argument into its own string
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error occurred when running command: %s %s\nError: %v\n", command, args, err)
	}
}
