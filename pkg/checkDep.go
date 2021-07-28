package JAA

import (
	"fmt"
	"os/exec"
	"runtime"
)

func CheckAll() bool {
	return (checkOS() && checkNmap())
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func checkOS() bool {
	// currently linux is supported.
	if runtime.GOOS != "linux" {
		return false
	}
	return true
}

func checkNmap() bool {
	if !commandExists("nmap") {
		fmt.Println("Nmap isn't installed")
		return false
	}
	return true
}

func checkVA() bool {
	return true
}
