package domain

import (
	crypt "vora/crypt"
	realm "vora/models/realm"
	rand "vora/utils/rand"
)

type Domain struct {
	id          string
	name        string
	totalRealms int
	realms      map[string]realm.Realm
}

func New(name string) Domain {
	var domain Domain
	domain.id = crypt.HashTo256(name + rand.String(15))
	domain.name = name
	domain.realms = map[string]realm.Realm{}
	return domain
}

func (d *Domain) AddRealm(name string) bool {
	if _, ok := d.realms[name]; ok {
		return false
	} else {
		d.realms[name] = realm.New(name)
		d.totalRealms += 1
		return true
	}

}

func (d *Domain) TotalRealms() int {
	return d.totalRealms
}

func (d *Domain) GetRealmByName(name string) realm.Realm {
	if _, ok := d.realms[name]; ok {
		return d.realms[name]
	} else {
		panic("realm not found")
	}
}

func (d *Domain) GetIdentity() string {
	return d.id
}

func (d *Domain) GetName() string {
	return d.name
}
