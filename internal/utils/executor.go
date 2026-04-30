package utils

import "os/exec"

func RunCmd(cmdStr string) (string, error) {
	cmd := exec.Command("bash", "-c", cmdStr)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
