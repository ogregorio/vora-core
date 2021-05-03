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
	prompt "graphgdb/prompt"

	cobra "github.com/spf13/cobra"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Prompt run graphgdb as command prompt",
	Long: `Prompt launches a graphgdb internal command prompt to speed up queries,
	 check commands and speed up interaction with the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt.Launch()
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}
