package main

import (
	"os"
	"os/exec"
	"runtime"
)

var operatingSystem = runtime.GOOS

func ClearScreen() {
	var cmd *exec.Cmd
	switch operatingSystem {
	case "darwin", "linux":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		panic("Your operating system is currently not supported, cannot clear the terminal screen")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
