package fixture

import "testing"

func TestFixture(t *testing.T) {
	if 1+1 != 2 {
		t.Fatal("unexpected")
	}
}
