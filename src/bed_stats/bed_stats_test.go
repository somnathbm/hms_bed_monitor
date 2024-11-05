package bed_stats

import (
	"testing"
)

func TestBedStats(t *testing.T) {
	got := GetBedStats()
	want := 45
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
