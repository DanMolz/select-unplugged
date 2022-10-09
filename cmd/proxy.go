package cmd

import (
	"github.com/neerolyte/select-unplugged/sp"
	"github.com/spf13/cobra"
)

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Share SP Pro serial via a proxy",
	Long:  `Share SP Pro serial via a proxy`,
	Run: func(cmd *cobra.Command, args []string) {
		proxy := new(sp.Proxy)
		proxy.Start()
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// proxyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// proxyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
