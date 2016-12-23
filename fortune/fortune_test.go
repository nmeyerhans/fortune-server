package fortune

import (
	"testing"
)

func TestAvailable(t *testing.T) {
	if !Available() {
		t.Error("fortune is not available")
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
}
