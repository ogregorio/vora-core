package commands

import (
	access "graphgdb/access"
)

func Exit() {

	access.UngrantAccess()

}
