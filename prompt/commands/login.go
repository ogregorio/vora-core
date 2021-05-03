package commands

import (
	errors "errors"
	access "graphgdb/access"
	gui "graphgdb/gui"
	errorTreat "graphgdb/io/error"

	promptui "github.com/manifoldco/promptui"
)

func Login() {

	username := user()
	password := password()

	access.Access(username, password)

	result, err := access.VerifyAccess()
	gui.P.Println(result)
	if err == nil && result {
		gui.G.Println("VERIFIED, you are logged.")
	} else {
		gui.R.Println("FAILED: " + err.Error())
	}

}

func user() string {
	validate := func(input string) error {
		if len(input) <= 4 {
			return errors.New("user must be longer than 4 characters")
		} else {
			return nil
		}
	}

	prompt := promptui.Prompt{
		Label:    "Username",
		Validate: validate,
	}

	result, err := prompt.Run()

	errorTreat.CheckError(err)

	return result
}

func password() string {
	validate := func(input string) error {
		if len(input) <= 4 {
			return errors.New("password must be longer than 4 characters")
		} else {
			return nil
		}
	}

	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()

	errorTreat.CheckError(err)

	return result
}
