package cmd

import (
	"fmt"

	"github.com/neerolyte/select-unplugged/sp"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of the SP Pro",
	Long:  `Logout of the SP Pro`,
	Run: func(cmd *cobra.Command, args []string) {
		spConnection := sp.ConnectionSerial{}
		spConnection.Open()
		protocol := sp.NewProtocol(&spConnection)
		err := protocol.Logout()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
