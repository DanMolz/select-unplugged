package cmd

import (
	"github.com/neerolyte/select-unplugged/sp"
	"github.com/spf13/cobra"
)

var fakeAddress string
var fakeCmd = &cobra.Command{
	Use:   "fake",
	Short: "Fake a SP Pro",
	Long:  `Fake a SP Pro`,
	Run: func(cmd *cobra.Command, args []string) {
		fake := new(sp.Fake)
		fake.Start(fakeAddress)
	},
}

func init() {
	rootCmd.AddCommand(fakeCmd)
	fakeCmd.Flags().StringVar(
		&fakeAddress,
		"address",
		"127.0.0.1:7528",
		"Address for fake inverter to listen on",
	)
}
