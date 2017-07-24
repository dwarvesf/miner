package commands

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	ConfigFileName = ".dfrc"
	BackupFileName = "dfrc.zip"
)

var (
	Out string
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

	subBackup.Flags().StringP("out", "o", "", "the back upfile directory")
	viper.BindPFlag("out", subBackup.Flags().Lookup("out"))
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
	path, err := getPathConfig()
	if err != nil {
		logrus.WithError(err).Error("cannot get path config file")
		return
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			logrus.WithError(err).Errorf("cannot create file config %s", ConfigFileName)
			return
		}

		logrus.Printf("%s successfully created, locates in ~/%s", ConfigFileName, ConfigFileName)
		return
	}

	logrus.Errorf("%s already existed", ConfigFileName)
}

func runBackupDotfiles() {
	path, err := getPathConfig()
	if err != nil {
		logrus.WithError(err).Error("cannot get path config file")
		return
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		logrus.WithError(err).Errorf("%s is not found", ConfigFileName)
		return
	}

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.WithError(err).Errorf("cannot read file config")
		return
	}

	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(bs)
	w.Close()

	bfile := BackupFileName
	if viper.Get("out").(string) != "" {
		bfile = viper.Get("out").(string)
	}

	err = ioutil.WriteFile(bfile, b.Bytes(), 0666)
	if err != nil {
		logrus.WithError(err).Errorf("cannot write file")
		return
	}
}

func runRestoreDotfiles() {
}

func runUpdateDotfiles() {
}

func runCleanupDotfiles() {
}

func getPathConfig() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", usr.HomeDir, ConfigFileName), nil
}
