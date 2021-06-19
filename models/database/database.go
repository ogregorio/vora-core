package database

import (
	domain "vora/models/domain"
)

type Database struct {
	totalDomains int
	domains      map[string]domain.Domain
}

// Instantiate generate a instance of a database
func (db *Database) Instantiate() {
	db.domains = make(map[string]domain.Domain)
}

// AddDomain return a bool if operation of add new Domains occurred with success or failed.
func (db *Database) AddDomain(name string) bool {
	domain := domain.New(name)
	db.domains[domain.GetIdentity()] = domain
	db.totalDomains++
	return true
}

// TotalDomains return a integer as number of domains registered.
func (db *Database) TotalDomains() int {
	return db.totalDomains
}

// GetDomainByName return a list of domains with passed name as string.
func (db *Database) GetDomainByName(name string) []domain.Domain {
	var domains []domain.Domain
	for _, element := range db.domains {
		if element.GetName() == name {
			domains = append(domains, element)
		}
	}
	return domains
}

// GetDomainById return a domain with passed id as string, or launch a panic if not exists.
func (db *Database) GetDomainById(id string) domain.Domain {
	if _, ok := db.domains[id]; ok {
		return db.domains[id]
	} else {
		panic("domain not found")
	}
}

// ListDomains returns a list of name of domains registered in database
func (db *Database) ListDomainsByName() []string {
	var domains []string
	for _, element := range db.domains {
		domains = append(domains, element.GetName())
	}
	return domains
}

// ListDomains returns a list of domains registered in database by id
func (db *Database) ListDomainsById() []string {
	var domains []string
	for key := range db.domains {
		domains = append(domains, key)
	}
	return domains
}
