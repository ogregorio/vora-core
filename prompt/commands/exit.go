package commands

import (
	access "vora-core/access"
)

func Exit() {

	access.UngrantAccess()

}
