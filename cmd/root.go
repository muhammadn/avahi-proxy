package cmd

import (
    "github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "avahi-proxy",
		Short: "Map Multicast DNS .local hostnames to your LAN domain name",
		Long: `Automatically map Multicast DNS .local hostnames to your LAN domain name (example, ".home.lan" or ".internal" and use your favourite DNS server to query avahi-proxy to get machine hostnames on .local multicast network. 

		Let your Windows and Android phones resolve your Multicast devices like iPhone/Mac and iPad, or maybe you want to expose .local hostnames over IPSec with your LAN domain name`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
