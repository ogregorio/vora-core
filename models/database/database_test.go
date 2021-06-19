package database

import (
	testing "testing"
)

var database Database

func TestDatabase(t *testing.T) {
	genDatabase()

	domains := database.ListDomainsByName()
	if len(domains) == 5 {
		t.Logf("Success, expected size: %d and received %d", 5, len(domains))
	} else {
		t.Errorf("Failed, expected size: %d and received %d", 5, len(domains))
	}

}

func TestDatabase2(t *testing.T) {
	genDatabase()

	domains := database.GetDomainByName("Test")
	if len(domains) == 3 {
		t.Logf("Success, expected size: %d and received %d", 3, len(domains))
	} else {
		t.Errorf("Failed, expected size: %d and received %d", 3, len(domains))
	}
}

func genDatabase() {
	database.Instantiate()
	database.AddDomain("Test")
	database.AddDomain("Test3")
	database.AddDomain("Test")
	database.AddDomain("Test2")
	database.AddDomain("Test")
}
