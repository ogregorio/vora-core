package commands

import (
	errors "errors"
	errorTreat "graphgdb/io/error"

	promptui "github.com/manifoldco/promptui"
)

func Username() string {
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

func Password() string {
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
