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
	"testing"
)

func TestAvailable(t *testing.T) {
	if !Available("/sbin/init") {
		t.Error("did not correctly identify file existence")
	}
	if Available("/sbin/xyzinit_xyz") {
		t.Error("did not identify file nonexistence")
	}
}

func TestFortuneExec(t *testing.T) {
	str, err := Fortune(true)
	if err != nil {
		t.Fatalf("Fortune returned unexpected error %s", err)
	}
	if len(str) == 0 {
		t.Error("Fortune returned successfully but with no content.")
	}
	str, err = Fortune(false)
	if err != nil {
		t.Fatalf("Fortune returned unexpected error %s", err)
	}
	if len(str) == 0 {
		t.Error("Fortune returned successfully but with no content.")
	}
}
