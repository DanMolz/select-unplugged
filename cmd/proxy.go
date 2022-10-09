package cmd

import (
	"github.com/neerolyte/select-unplugged/sp"
	"github.com/spf13/cobra"
)

var proxyAddress string
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Share SP Pro serial via a proxy",
	Long:  `Share SP Pro serial via a proxy`,
	Run: func(cmd *cobra.Command, args []string) {
		proxy := new(sp.Proxy)
		proxy.Start(proxyAddress)
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)
	proxyCmd.Flags().StringVar(
		&proxyAddress,
		"address",
		"127.0.0.1:7528",
		"Address for proxy to listen on",
	)
}
