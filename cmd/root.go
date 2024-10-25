package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "avahi-proxy",
    Short: "Avahi Proxy allows you to register Multicast DNS devices to your normal DNS",
    Long:  "Avahi Proxy is a daemon that allows you to register Multicast DNS devices by acting as a proxy from your standard DNS server to Multicast DNS records.",
    Run: func(cmd *cobra.Command, args []string) {

    },
}
