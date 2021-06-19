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
	"fmt"

	"github.com/spf13/cobra"
)

var databaseName string

// databaseCmd represents the database command
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "High-level database manipulate commands.",
	Long:  `High-level database manipulate commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("database called")
	},
}

func init() {
	databaseCmd.PersistentFlags().StringVarP(&databaseName, "name", "n", "", "create a database name")
	databaseCmd.MarkFlagRequired("name")
	databaseCmd.AddCommand(addCmd)
	rootCmd.AddCommand(databaseCmd)
}
