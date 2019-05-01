package cmd

import (
	"log"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// nolint: gochecknoglobals
	cfgFile string
)

func setupOtherCmd(rootCmd *cobra.Command) {
	var msgOtherCmd = &cobra.Command{
		Use:   "othermsg",
		Short: "othermsg",
		Long:  `My other message`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// Get root log flag
			l := viper.GetString("log")

			// Get msg flag
			msg := viper.GetString("msg")
			return sendOtherMsg(msg, l)
		},
	}
	rootCmd.AddCommand(msgOtherCmd)
	msgOtherCmd.Flags().String("msg", "And now for something completely di...", "msg to send")
	if err := viper.BindPFlag("msg", msgOtherCmd.Flags().Lookup("msg")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}
}

func setupCmd(rootCmd *cobra.Command) {
	var msgCmd = &cobra.Command{
		Use:   "msg",
		Short: "msg",
		Long:  `My message`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// Get root log flag
			l := viper.GetString("log")

			// Get msg flag
			msg := viper.GetString("msg")
			return sendMsg(msg, l)
		},
	}
	rootCmd.AddCommand(msgCmd)
	msgCmd.Flags().String("msg", "We are the Knights who say.....   \"Ni\"!", "msg to send")
	if err := viper.BindPFlag("msg", msgCmd.Flags().Lookup("msg")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}
}

func Execute() {

	// Setup global root command
	var rootCmd = &cobra.Command{
		Use:   "root",
		Short: "root command",
		Long:  `root command`,
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().String("log", "knights.of.ni", "Syslog host")
	if err := viper.BindPFlag("log", rootCmd.PersistentFlags().Lookup("log")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}

	setupOtherCmd(rootCmd)
	setupCmd(rootCmd)

	cobra.OnInitialize(initConfig)
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
