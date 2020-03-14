package radCor

import "testing"

func TestHello(t *testing.T) {
	want := "it is ok"
	if got := Coin(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
