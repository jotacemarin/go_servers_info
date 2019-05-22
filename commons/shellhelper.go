package commons

import (
	"fmt"
	"os/exec"
	"runtime"
)

// ShellCall : func
func ShellCall(command string, parameter string) (map[string]string, error) {
	if thismachine := runtime.GOOS; thismachine == "windows" {
		return nil, fmt.Errorf("Cant execute this command in %s", thismachine)
	}
	ser, err := shellExecution(command, parameter)
	if err != nil {
		return nil, err
	}
	return ser, nil
}

// ShellExecution : func
func shellExecution(command string, parameter string) (map[string]string, error) {
	out, err := exec.Command(command, parameter).Output()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	output := string(out[:])
	ser := make(map[string]string)
	ser["output"] = output
	return ser, nil
}
