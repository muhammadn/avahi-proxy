package cmd

import (
        "github.com/spf13/cobra"
	"muhammadn/avahi-proxy/pkg/proxy"
)

var runCmd = &cobra.Command{
    Use:   "run",
    Aliases: []string{"r"},
    Short:  "Start running the proxy",
    RunE: func(cmd *cobra.Command, args []string) error {
        baseDomain, _               := cmd.Flags().GetString("baseDomain")
	port, _                     := cmd.Flags().GetString("port")

	baseDomain = baseDomain + "."

	proxy.RunProxy(baseDomain, port)
        return nil
    },
}

func init() {
        rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("baseDomain", "", "", "Your local LAN domain name, example: home.lan")
	runCmd.Flags().StringP("port", "p", "5354", "Port number to run avahi-proxy (default: 5354)")

        runCmd.MarkFlagRequired("baseDomain")
}
