package cmd

import (
	"log"

	"github.com/neerolyte/select-unplugged/sp"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout to the SP Pro",
	Long:  `Logout to the SP Pro`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Connecting to SP Pro")
		spConnection := sp.ConnectionSerial{}
		spConnection.Open()
		protocol := sp.NewProtocol(&spConnection)
		memory := sp.NewMemory(40973 /* SP Link Disconnecting Comms 1 */, 1)
		memory.SetData([]byte("\x01\x00"))
		request := sp.NewRequestWrite(memory)
		response, _ := protocol.Send(request)
		log.Printf("Response: %s", response)
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
