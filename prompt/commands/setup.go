package commands

import (
	"strings"
	access "vora-core/access"
	crypt "vora-core/crypt"
	gui "vora-core/gui"
	files "vora-core/utils/files"

	viper "github.com/spf13/viper"
	babble "github.com/tjarratt/babble"
)

func Setup() {

	// create generatrix
	generatrix()

	// create access for user
	createAccess("Welcome to setup!", 0)

}

func createAccess(message string, count byte) {
	gui.P.Println(message)

	username := Username()
	password := Password()
	confirmPassword := Password()

	// TODO: Retry when count pass 3
	if confirmPassword != password && count < 3 {
		createAccess("Ops, password don't match, try again!", count+1)
	}

	access.CreateAccess(username, password)

}

func generatrix() {
	gui.P.Println("Write this passphrases (this is the last time you will see they here, but never forget! ) ")
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
