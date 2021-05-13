package run

import (
	strings "strings"
	access "vora/access"
	crypt "vora/crypt"
	gui "vora/gui"
	files "vora/utils/files"

	viper "github.com/spf13/viper"
	babble "github.com/tjarratt/babble"
)

// Setup start a initial setup process
func Setup(login string, password string) {

	// create generatrix
	generatrix()

	// create access for user
	createAccess(login, password)

}

// createAccess provides a simple access to access package
func createAccess(login string, password string) {
	access.CreateAccess(login, password)
}

// generatrix gen a generatrix for system
func generatrix() {
	gui.P.Print("Write this passphrases")
	gui.R.Println(" (this is the last time you will see they here, never forget!)")
	babbler := babble.NewBabbler()
	babbler.Count = 12
	passphrase := babbler.Babble()
	oneByOnePhrases := strings.Split(passphrase, "-")
	gui.Y.Println("Your exclusive passphrases: ")
	for i := 0; i < len(oneByOnePhrases); i++ {
		gui.Y.Printf("[ %2.d ] %s \n", i+1, oneByOnePhrases[i])
	}
	gui.P.Println("Generate yor generatrix...")
	files.Write(viper.GetString("root"), string(crypt.Keygen(passphrase)))
}
