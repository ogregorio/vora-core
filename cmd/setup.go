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
	commands "vora-core/prompt/commands"

	cobra "github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup vora-core in yout machine.",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Setup()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
