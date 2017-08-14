// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// 	You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
