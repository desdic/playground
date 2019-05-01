package cmd

import (
	"fmt"
)

func sendMsg(msg, loghost string) error {
	fmt.Printf("My message is '%s' and I'm logging it to '%s'\n", msg, loghost)
	return nil
}
