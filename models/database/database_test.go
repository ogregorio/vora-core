package database

import (
	testing "testing"
)

func TestDatabase(t *testing.T) {
	var database Database
	database.Instantiate()
	database.AddDomain("Test")
	domains := database.ListDomains()
	for _, domain := range domains {
		t.Logf(domain)
	}
}
