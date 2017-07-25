package commands

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
)

func cleanUpLanguages(langs []string) {
	for _, v := range langs {
		if err := brewUninstall(v); err != nil {
			logrus.WithError(err).Errorf("failed to install language %s", v)
		}
	}
}

func cleanUpTools(tools []string) {
	for _, v := range tools {
		if err := brewUninstall(v); err != nil {
			logrus.WithError(err).Errorf("failed to install tool %s", v)
		}
	}
}

func cleanUpEditors(editors []string) {
	for _, v := range editors {
		if err := brewUninstall(v); err != nil {
			logrus.WithError(err).Errorf("failed to install editor %s", v)
		}
	}
}

func brewUninstall(name string) error {
	cmd := exec.Command("brew", "uninstall", name)
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))

	return cmd.Run()
}
