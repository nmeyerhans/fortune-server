package fortune

import (
	"encoding/base64"
	"os"
	"os/exec"
)

const (
	FortuneCmd = "/usr/games/fortune"
)

func Available() (bool) {
	_, err := os.Stat(FortuneCmd)
	if err != nil {
		return false
	}
	return true
}

func Fortune(formatJSON bool) (string, error) {
	cmd := exec.Command(FortuneCmd)
	output, err := cmd.CombinedOutput()
	if(err != nil) {
		return "", err
	}
	if formatJSON {
		return base64.StdEncoding.EncodeToString(output), nil
	} else {
		return string(output), nil
	}
}
