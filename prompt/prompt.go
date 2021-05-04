package prompt

import (
	"graphgdb/gui"
	commands "graphgdb/prompt/commands"

	promptui "github.com/manifoldco/promptui"
)

func Launch() {
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "graphgdb >> ",
		Validate: validate,
	}

	var err error
	var result string

	for err == nil {
		result, err = prompt.Run()
		call(result)
	}

	gui.P.Println("last command result: " + result)

}

func call(input string) {
	switch input {
	case "login":
		commands.Login()
	case "exit":
		commands.Exit()
	}
}
