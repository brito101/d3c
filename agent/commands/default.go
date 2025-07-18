package commands

import (
	"os/exec"
	"runtime"
)

type Default struct {
	Command string
}

func (instance Default) Exec() (response string, err error) {
	switch runtime.GOOS {
	case "windows":
		output, err := exec.Command("powershell.exe", "/C", instance.Command).CombinedOutput()
		response = string(output)
		return response, err
	case "linux":
		output, err := exec.Command("bash", "-c", instance.Command).Output()
		response = string(output)
		return response, err
	default:
		response = "System not implemented: " + runtime.GOOS
		return response, nil
	}
}
