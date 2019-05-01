package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

func init() {
	rootCmd.AddCommand(msgOtherCmd)

	msgOtherCmd.Flags().String("msg", "And now for something completely di...", "msg to send")
	if err := viper.BindPFlag("msg", msgOtherCmd.Flags().Lookup("msg")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}
}
func sendOtherMsg(msg, loghost string) error {
	fmt.Printf("My message is '%s' and I'm logging it to '%s'\n", msg, loghost)
	return nil
}
