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
	cobra "github.com/spf13/cobra"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor [OPTIONS] [COMMANDS]",
	Short: "Run a doctor to validate the components or events situation",
	Long: `Run a doctor to validate the components or events situation,
	this will ensure that your database is running in good health`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "all":
			//doctor.doctor()
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
