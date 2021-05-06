package commands

import (
	errors "errors"
	errorTreat "graphgdb/io/error"

	promptui "github.com/manifoldco/promptui"
)

func Default() string {
	validate := func(input string) error {
		if len(input) <= 1 {
			return errors.New("ops, command not recognized")
		} else {
			return nil
		}
	}

	prompt := promptui.Prompt{
		Label:    "graphgdb",
		Validate: validate,
	}

	result, err := prompt.Run()

	errorTreat.CheckError(err)

	return result
}
