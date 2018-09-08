package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version string = "v0.0.1"
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "show",
	Short: "Git Repositry Description getter",
	Long:  "Git Repositry Description getter",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	for _, arg := range os.Args {
		if arg == "--version" || arg == "-v" {
			fmt.Println("grd " + version)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("version", "v", false, "print grd version")

	rootCmd.AddCommand(versionCmd)

	viper.SetDefault("author", "taxio")
	viper.SetDefault("licecse", "MIT")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".tmp")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of grd",
	Long:  "This is grd",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
