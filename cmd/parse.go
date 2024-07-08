package cmd

import (
	"encoding/hex"
	"fmt"
	"log"
	"regexp"

	"github.com/neerolyte/select-unplugged/sp"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse <message>",
	Short: "Parse a message",
	Long:  `Parse a message`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s := args[0]
		s = regexp.MustCompile("^0x").ReplaceAllLiteralString(s, "")
		data, err := hex.DecodeString(s)
		if err != nil {
			panic(err)
		}
		message := sp.Message(data)
		fmt.Println(message.Describe())
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
