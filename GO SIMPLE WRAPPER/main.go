package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Path to your executable
	cmd := exec.Command("C:/PATH/TO/EXCECUTABLE.EXE")

	// Suppress the command window
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	// Execute the command and get the output
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print the output
	fmt.Println(string(out))
}
