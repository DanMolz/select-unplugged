package cmd

import (
	"fmt"

	"github.com/neerolyte/select-unplugged/sp"
	"github.com/spf13/cobra"
)

var statSelectPassword string
var statSelectCmd = &cobra.Command{
	Use:   "stat-select",
	Short: "select.live device emulation CLI",
	Long:  `select.live device emulation CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		// protocol := NewConnectedProtocol()

		spConnection := sp.NewConnectionSerial(SerialPort)
		spConnection.Open()
		protocol := sp.NewProtocol(&spConnection)
		err := protocol.Login(loginPassword)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Success\n")

		fmt.Print(sp.StatsSelectLiveRender(protocol))
	},
}

func init() {
	rootCmd.AddCommand(statSelectCmd)
}
