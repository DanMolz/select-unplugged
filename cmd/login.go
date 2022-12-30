package cmd

import (
	"log"

	"github.com/neerolyte/select-unplugged/sp"
	"github.com/spf13/cobra"
)

var loginPassword string
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the SP Pro",
	Long:  `Login to the SP Pro`,
	Run: func(cmd *cobra.Command, args []string) {
		loginHashAddress := sp.Address(2031616)
		loginHashWords := sp.Words(8)
		loginHashArea := sp.NewArea(loginHashAddress, loginHashWords)
		loginHashMemory, _ := sp.NewMemory(loginHashAddress, loginHashWords)
		loginStatusAddress := sp.Address(0x1f0010)
		loginStatusWords := sp.Words(1)
		loginStatusArea := sp.NewArea(loginStatusAddress, loginStatusWords)
		log.Printf("Connecting to SP Pro")
		spConnection := sp.ConnectionSerial{}
		spConnection.Open()
		protocol := sp.NewProtocol(&spConnection)

		log.Printf("Requesting login hash")
		readLoginHashRequest := sp.NewRequestQuery(loginHashArea)
		loginHashResponse, _ := protocol.Send(readLoginHashRequest)
		loginHash, _ := sp.Message(loginHashResponse).Data()
		log.Printf("Login hash: %x", *loginHash)
		log.Printf("Password: %s", loginPassword)

		responseHash := sp.CalculateLoginHash(loginPassword, *loginHash)

		loginHashMemory.SetData(sp.Data(responseHash))
		writeLoginHashRequest := sp.NewRequestWrite(*loginHashMemory)
		loginWriteHashResponse, _ := protocol.Send(writeLoginHashRequest)
		log.Printf("Login hash: %s", &loginWriteHashResponse)

		log.Printf("Requesting login status")
		readLoginStatusRequest := sp.NewRequestQuery(loginStatusArea)
		loginStatusResponse, _ := protocol.Send(readLoginStatusRequest)
		log.Printf("Login status: %s", &loginStatusResponse)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVar(
		&loginPassword,
		"password",
		"Selectronic SP Pro",
		"Password to login with",
	)
}
