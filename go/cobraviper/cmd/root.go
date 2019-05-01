package cmd

import (
	"log"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root command",
	Long:  `root command`,
}

func Execute() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().String("log", "knights.of.ni", "Syslog host")
	if err := viper.BindPFlag("log", rootCmd.PersistentFlags().Lookup("log")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(home)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if cfgFile != "" {
			log.Println("config specified but unable to read it, using defaults")
		}
	}
}
