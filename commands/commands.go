package commands

import (
	"fmt"
	"os"

	"github.com/dwarvesf/miner/config"
	"github.com/dwarvesf/miner/version"
	"github.com/spf13/cobra"
)

var (
	conf *config.Config
)

func New() *cobra.Command {
	cobra.OnInitialize()
	conf = config.NewConfig("")

	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "miner",
		Short: "The mining assistant of Dwarves Foundation",
		Long:  `Miner comes along with the dwarves, show his incredible magic when dwarves need. Without miner, Dwarves Foundation cannot go far this long.`,
	}

	dotfilesCmd := &cobra.Command{
		Use:   "dotfiles",
		Short: "Dotfiles commands",
		Run: func(cmd *cobra.Command, args []string) {
			runDotfiles()
			os.Exit(0)
		},
	}
	rootCmd.AddCommand(dotfilesCmd)

	versionCommand := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of miner",
		Long:  `All software has versions. This is miner's.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Miner's version %s", version.VERSION)
		},
	}
	rootCmd.AddCommand(versionCommand)

	return rootCmd
}

func runDotfiles() {

}
