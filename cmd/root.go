package cmd

import (
	bytes "bytes"
	fmt "fmt"
	errorTreat "graphgdb/io/error"
	files "graphgdb/utils/files"
	os "os"

	cobra "github.com/spf13/cobra"
	viper "github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "graphgdb",
	Short: "GraphgDB is a fast and simple database graph based.",
	Long:  `A fast and flexible database graph based`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func initConfig() {
	viper.SetConfigName("config")          // name of config file
	viper.SetConfigType("yaml")            // type of config file
	viper.AddConfigPath("$HOME/graphgdb/") // where is file

	// try to get confi file
	err := viper.ReadInConfig()
	errorTreat.CheckError(err)

	// get content from file
	config, err := files.ReadAllBytes(viper.ConfigFileUsed())
	errorTreat.CheckError(err)

	// get configs from content
	viper.ReadConfig(bytes.NewBuffer(config))
}
