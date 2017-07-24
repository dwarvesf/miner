package commands

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	ConfigFileName = ".dfrc"
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

// runInitDotfiles creates .dfrc if not exist
func runInitDotfiles() {
	usr, err := user.Current()
	if err != nil {
		logrus.WithError(err).Error("cannot get current user")
		return
	}

	path := fmt.Sprintf("%s/%s", usr.HomeDir, ConfigFileName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			logrus.WithError(err).Errorf("cannot create file config %s", ConfigFileName)
			return
		}

		logrus.Printf("%s successfully created, locates in %s/%s", ConfigFileName, usr.HomeDir, ConfigFileName)
		return
	}

	logrus.Errorf("%s already existed", ConfigFileName)
}

func runBackupDotfiles() {
}

func runRestoreDotfiles() {
}

func runUpdateDotfiles() {
}

func runCleanupDotfiles() {
}
