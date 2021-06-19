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
	cmdstructs "vora/cmd/cmdstructs"
	run "vora/run"

	cobra "github.com/spf13/cobra"
)

var add cmdstructs.Add

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Command to add any type of elements",
	Long:  `Command to add any type of elements`,
	Run: func(cmd *cobra.Command, args []string) {
		run.Add(add)
	},
}

func init() {
	addCmd.PersistentFlags().StringVarP(&add.Domain, "domain", "", "", "select domain to work")
	addCmd.PersistentFlags().StringVarP(&add.Model, "model", "m", "", "declare a type of model into the domain")
	addCmd.PersistentFlags().StringVarP(&add.Name, "name", "n", "", "declare a name of model into the domain")
	addCmd.PersistentFlags().StringVarP(&add.Source, "source", "s", "", "declare a source of model into the domain")
	addCmd.PersistentFlags().StringVarP(&add.Destiny, "destiny", "d", "", "declare a destiny of model into the domain")
	rootCmd.AddCommand(addCmd)
}
