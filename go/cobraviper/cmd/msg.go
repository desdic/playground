package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

func init() {
	rootCmd.AddCommand(msgCmd)

	msgCmd.Flags().String("msg", "We are the Knights who say.....   \"Ni\"!", "msg to send")
	if err := viper.BindPFlag("msg", msgCmd.Flags().Lookup("msg")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}
}
func sendMsg(msg, loghost string) error {
	fmt.Printf("My message is '%s' and I'm logging it to '%s'\n", msg, loghost)
	return nil
}
