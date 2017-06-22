package commands

import (
	"fmt"

	"github.com/goguard/goguard/version"
	"github.com/spf13/cobra"
)

// VERSION ...
const (
	VERSION = "1.0"
)

func initVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of miner",
		Long:  `All software has versions. This is miner's.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Miner's version %s", version.VERSION)
		},
	}
}
