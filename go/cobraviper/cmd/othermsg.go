package cmd

import (
	"fmt"
)

func sendOtherMsg(msg, loghost string) error {
	fmt.Printf("My message is '%s' and I'm logging it to '%s'\n", msg, loghost)
	return nil
}
