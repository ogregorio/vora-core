package run

import (
	os "os"
	access "vora/access"
	gui "vora/gui"
)

func Login(login string, password string) {
	loginGrantPermissions(login, password)

}

func loginGrantPermissions(login string, password string) bool {
	err := access.GrantAccess(login, password)
	if err == nil {
		gui.DMessage("PASS: you are logged", "G")
		return true
	} else {
		gui.DMessage("NOT PASS: "+err.Error(), "R")
		os.Exit(-1)
		return false
	}
}
