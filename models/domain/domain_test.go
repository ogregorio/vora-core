package domain

import (
	testing "testing"
)

func TestDomain(t *testing.T) {
	var domain Domain

	d.New("Test")
	d.AddRealm("Teste")

	if d.TotalRealms() == 1 {
		t.Logf("Success! Expected %d and received %d", 1, d.TotalRealms())
	} else {
		t.Errorf("Failed! Expected %d and received %d", 1, d.TotalRealms())
	}
}
