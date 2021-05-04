package commands

import (
	access "graphgdb/access"
	gui "graphgdb/gui"
)

func Login() {

	username := Username()
	password := Password()

	err := access.GrantAccess(username, password)

	if err == nil {
		gui.G.Println("PASS: Permissions granted.")
	} else {
		gui.R.Println("FAILED: " + err.Error())
	}
}
