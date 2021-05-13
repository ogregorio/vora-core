package commands

import (
	access "vora-core/access"
	gui "vora-core/gui"
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
