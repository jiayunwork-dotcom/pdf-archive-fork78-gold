package config

import "testing"

func TestDefaultConfig_LowConfThreshold(t *testing.T) {
	got := DefaultConfig().Pipeline.LowConfThreshold
	if got != 0.6 {
		t.Fatalf("DefaultConfig LowConfThreshold=%v, want 0.6", got)
	}
}
