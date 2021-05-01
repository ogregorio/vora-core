package cli

import (
	constants "graphgdb/constants"
)

func Test(argv ...string) {
	ConstantsTest()
}

func ConstantsTest() {
	result, err := constants.ConstantsTest()

	if result && err == nil {
		G.Println("Constants import tool : PASS")
	} else {
		R.Println("Constants import tool : NOT PASS\n" + err.Error())
	}
}
