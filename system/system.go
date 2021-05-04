package system

var logged bool
var initSystem bool

func System() bool {
	initSystem = true
	return initSystem
}

func IsLogged() bool {
	return logged
}

func SetLogged(status bool) {
	logged = status
}
