package commons

import (
	"fmt"
	"os/exec"
	"runtime"
)

// ShellCall : func
func ShellCall(command string, parameter string) (string, error) {
	if thismachine := runtime.GOOS; thismachine == "windows" {
		return "", fmt.Errorf("Cant execute this command in %s", thismachine)
	}
	ser, err := shellExecution(command, parameter)
	if err != nil {
		return "", err
	}
	return ser, nil
}

// ShellExecution : func
func shellExecution(command string, parameter string) (string, error) {
	out, err := exec.Command(command, parameter).Output()
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	output := string(out[:])
	return output, nil
}
