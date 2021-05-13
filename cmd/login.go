/*
Copyright © 2020 Arthur Gregório <arthurgregorioleal@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	access "vora/access"
	gui "vora/gui"

	cobra "github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login provides a way to guarantee access to data and database files.",
	Long: `Login provides a way to guarantee access to data and database files.
	It expects a username and password entered in the initial setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "verify":
			loginCmdVerify()
		default:
			loginGrantPermissions(args[0], args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func loginCmdVerify() {
	result := access.VerifyAccess()
	if result {
		gui.G.Println("PASS: you are logged.")
	} else {
		gui.R.Println("FAILED: you aren't logged.")
	}
}

func loginGrantPermissions(login string, password string) {
	err := access.GrantAccess(login, password)
	if err == nil {
		gui.G.Println("PASS: Permissions granted.")
	} else {
		gui.R.Println("FAILED: " + err.Error())
	}
}
