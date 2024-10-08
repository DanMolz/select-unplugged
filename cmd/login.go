package cmd

import (
	"fmt"

	"github.com/neerolyte/select-unplugged/sp"
	"github.com/spf13/cobra"
)

var loginPassword string
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the SP Pro",
	Long:  `Login to the SP Pro`,
	Run: func(cmd *cobra.Command, args []string) {
		spConnection := sp.NewConnectionSerial(SerialPort)
		spConnection.Open()
		protocol := sp.NewProtocol(&spConnection)
		err := protocol.Login(loginPassword)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Success\n")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVar(
		&loginPassword,
		"password",
		"Selectronic SP PRO",
		"Password to login with",
	)
}
