package cmd

import (
	"time"

	"github.com/neerolyte/select-unplugged/sp"
	"github.com/spf13/cobra"
)

var statSelectCmd = &cobra.Command{
	Use:   "stat-select",
	Short: "select.live device emulation CLI",
	Long:  `select.live device emulation CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		spConnection := sp.NewConnectionSerial(SerialPort)
		spConnection.Open()
		protocol := sp.NewProtocol(&spConnection)
		err := protocol.Login(loginPassword)
		if err != nil {
			panic(err)
		}

		// Initial execution of StatsSelectLiveRenderV2
		sp.StatsSelectLiveRenderV2(protocol)

		// Create a channel to wait indefinitely
		done := make(chan struct{})
		defer close(done)

		// Run StatsSelectLiveRenderV2 every 30 seconds in a goroutine
		go func() {
			ticker := time.NewTicker(15 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					sp.StatsSelectLiveRenderV2(protocol)
				case <-done:
					return
				}
			}
		}()

		// Wait indefinitely
		<-done
	},
}

func init() {
	rootCmd.AddCommand(statSelectCmd)
}
