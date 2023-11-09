package mascot_test

import (
	"testing"

	"example.com/go-demo-1/mascot"
)

func TestBestMascot(t *testing.T) {
	want := "TChalla"
	if got := mascot.BestMascot(); got != want {
		t.Errorf("BestMascot() = %q, want %q", got, want)
	}
}
