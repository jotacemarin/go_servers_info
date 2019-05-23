package commons

import (
	"fmt"
	"log"
	"runtime"

	pipes "github.com/ebuchman/go-shell-pipes"
)

// ShellCall : func
func ShellCall(command string, parameter string, grep string) (string, error) {
	if thismachine := runtime.GOOS; thismachine == "windows" {
		return "", fmt.Errorf("Cant execute this command in %s", thismachine)
	}
	ser, err := shellExecution(command, parameter, grep)
	if err != nil {
		return "", err
	}
	return ser, nil
}

// ShellExecution : func
func shellExecution(command string, parameter string, grep string) (string, error) {
	s, err := pipes.RunString(fmt.Sprintf("%s %s | grep %s", command, parameter, grep))
	if err != nil {
		log.Print(err)
	}
	return s, nil
}
