package utils

import "testing"

func TestNormalizeDate_ChineseFullDate(t *testing.T) {
	got := NormalizeDate("2024年3月15日")
	if got != "2024-03" {
		t.Fatalf("NormalizeDate(%q)=%q, want 2024-03", "2024年3月15日", got)
	}
}
