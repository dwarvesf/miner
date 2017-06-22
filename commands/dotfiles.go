package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func initDotfilesCommand() *cobra.Command {
	dotfilesCmd := &cobra.Command{
		Use:   "dotfiles",
		Short: "dotfiles commands",
	}

	subInit := &cobra.Command{
		Use:   "init dotfiles",
		Short: "dotfiles initialization",
		Run: func(cmd *cobra.Command, args []string) {
			runInitDotfiles()
			os.Exit(0)
		},
	}
	dotfilesCmd.AddCommand(subInit)

	subBackup := &cobra.Command{
		Use:   "backup dotfiles",
		Short: "dotfiles initialization",
		Run: func(cmd *cobra.Command, args []string) {
			runBackupDotfiles()
			os.Exit(0)
		},
	}
	dotfilesCmd.AddCommand(subBackup)

	subRestore := &cobra.Command{
		Use:   "restore dotfiles",
		Short: "restore dotfiles initialization",
		Run: func(cmd *cobra.Command, args []string) {
			runRestoreDotfiles()
			os.Exit(0)
		},
	}
	dotfilesCmd.AddCommand(subRestore)

	subUpdate := &cobra.Command{
		Use:   "update dotfiles",
		Short: "update dotfiles initialization",
		Run: func(cmd *cobra.Command, args []string) {
			runUpdateDotfiles()
			os.Exit(0)
		},
	}
	dotfilesCmd.AddCommand(subUpdate)

	subCleanup := &cobra.Command{
		Use:   "cleanup dotfiles",
		Short: "clean up dotfiles",
		Run: func(cmd *cobra.Command, args []string) {
			runCleanupDotfiles()
			os.Exit(0)
		},
	}
	dotfilesCmd.AddCommand(subCleanup)

	return dotfilesCmd
}

func runInitDotfiles() {
}

func runBackupDotfiles() {
}

func runRestoreDotfiles() {
}

func runUpdateDotfiles() {
}

func runCleanupDotfiles() {
}
