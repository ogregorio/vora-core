package models

import realm "vora/models/realm"

var realms []realm.Realm

func NewDomain() {
	realms = []realm.Realm{}
}

func AddNode(r string, n string) {
	gr := GetRealmByName(r)
	realm.AddNode(gr, n)
}

func AddEdge(r string, s1 string, s2 string) {
	gr := GetRealmByName(r)
	realm.AddEdge(gr, s1, s2)
}

func Addrealm(name string) {
	addRealm(realm.New(name))
}

func addRealm(r realm.Realm) {
	realms = append(realms, r)
}

func TotalRealms() int {
	return len(realms)
}

func GetRealmByName(name string) realm.Realm {
	for i := 0; i < len(realms); i++ {
		if realm.GetName(realms[i]) == name {
			return realms[i]
		}
	}
	panic("realm not found")
}
