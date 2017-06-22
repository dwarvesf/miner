package commands

import (
	"github.com/dwarvesf/miner/config"
	"github.com/spf13/cobra"
)

var (
	conf *config.Config
)

// New ...
func New() *cobra.Command {
	cobra.OnInitialize()
	conf = config.NewConfig("")

	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "miner",
		Short: "The mining assistant of Dwarves Foundation",
		Long:  `Miner comes along with the dwarves, show his incredible magic when dwarves need. Without miner, Dwarves Foundation cannot go far this long.`,
	}

	rootCmd.AddCommand(initDotfilesCommand())
	rootCmd.AddCommand(initVersionCommand())
	rootCmd.AddCommand(initEmailCommand())
	rootCmd.AddCommand(initProjectCommand())

	return rootCmd
}
