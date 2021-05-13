package cmd

import (
	bytes "bytes"
	fmt "fmt"
	os "os"
	strings "strings"
	gui "vora/gui"
	errorTreat "vora/io/error"
	run "vora/run"
	files "vora/utils/files"

	cobra "github.com/spf13/cobra"
	viper "github.com/spf13/viper"
)

var (
	connect string
	setup   []string
	debug   bool
)

var rootCmd = &cobra.Command{
	Use:   "vora",
	Short: "Vora is a consistent and simple database graph based.",
	Long:  `A consistent and flexible database graph based.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		preRun()
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	initConfig()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}

func initConfig() {

	rootCmd.PersistentFlags().StringVarP(&connect, "connect", "c", "", "connect using user and pass with user@password")
	rootCmd.PersistentFlags().StringSliceVarP(&setup, "setup", "s", nil, "start initial setup process: true user@password")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show debug messages during execution")

	viper.SetConfigName("config")      // name of config file
	viper.SetConfigType("yaml")        // type of config file
	viper.AddConfigPath("$HOME/vora/") // where is file

	// try to get confi file
	err := viper.ReadInConfig()
	errorTreat.CheckError(err)

	// get content from file
	config, err := files.ReadAllBytes(viper.ConfigFileUsed())
	errorTreat.CheckError(err)

	// get configs from content
	viper.ReadConfig(bytes.NewBuffer(config))
}

func preRun() {
	if debug {
		gui.DEBUG = true
	} else {
		gui.DEBUG = false
	}

	if setup != nil {
		if setup[0] == "true" {
			temp := strings.Split(setup[1], "@")
			run.Setup(temp[0], temp[1])
		}
	} else {
		temp := strings.Split(connect, "@")
		run.Login(temp[0], temp[1])
	}
}
