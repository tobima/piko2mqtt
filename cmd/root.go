package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tobima/piko2mqtt/pkg/dxs"
	"github.com/tobima/piko2mqtt/pkg/piko"
)

var (
	// Used for flags.
	host string

	rootCmd = &cobra.Command{
		Use:   "piko2mqtt",
		Short: "piko2mqtt - connects Kostal Piko inverter to MQTT",
		Long:  "",

		Run: func(cmd *cobra.Command, args []string) {
			inv := piko.NewInverter(host)
			data := inv.GetTypePlate()
			dxs.PrintEntries(data.DxsEntries)
			data = inv.GatherAC()
			dxs.PrintEntries(data.DxsEntries)

			fmt.Printf("MPPT1 %s\n", inv.GatherMPPT(1))
			fmt.Printf("MPPT2 %s\n", inv.GatherMPPT(2))
			fmt.Printf("MPPT3 %s\n", inv.GatherMPPT(3))
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&host, "host", "192.168.178.185", "IP or hostname of the inverter")

}

func initConfig() {
	// add code to load from config file with viper
}
