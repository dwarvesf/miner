package commands

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
)

func installLanguages(langs []string) {
	for _, v := range langs {
		if err := brewInstall(v); err != nil {
			logrus.WithError(err).Errorf("failed to install language %s", v)
		}
	}
}

func installTools(tools []string) {
	for _, v := range tools {
		if err := brewInstall(v); err != nil {
			logrus.WithError(err).Errorf("failed to install tool %s", v)
		}
	}
}

func installEditors(editors []string) {
	for _, v := range editors {
		if err := brewInstall(v); err != nil {
			logrus.WithError(err).Errorf("failed to install editor %s", v)
		}
	}
}

func brewInstall(name string) error {
	cmd := exec.Command("brew", "install", name)
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))

	return cmd.Run()
}
