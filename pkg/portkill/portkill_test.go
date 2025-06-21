package portkill

import "testing"

func TestKillPort(t *testing.T) {
	err := KillPort(5001)
	if err != nil {
		t.Errorf("kill port failed: %w", err)
	}
}
