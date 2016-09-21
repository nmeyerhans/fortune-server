package fortune

import (
	"encoding/base64"
	"os/exec"
)

func Fortune(formatJSON bool) (string, error) {
	cmd := exec.Command("fortune")
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
