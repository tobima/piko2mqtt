package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tobima/piko2mqtt/pkg/piko"
)

var rootCmd = &cobra.Command{
	Use:   "piko2mqtt",
	Short: "piko2mqtt - connects Kostal Piko inverter to MQTT",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		inv := piko.NewInverter("192.168.178.185")
		data := inv.GetTypePlate()

		for _, entry := range data.DxsEntries {
			fmt.Printf("%s: %s \n", entry.ID, entry.Value)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
