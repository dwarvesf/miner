package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func initProjectCommand() *cobra.Command {
	projectCmd := &cobra.Command{
		Use:   "project",
		Short: "projects commands",
	}

	subNew := &cobra.Command{
		Use:   "new",
		Short: "project initialization",
		Long:  "setup gitlab group, repositories with template, trello board, slack channel",
		Run: func(cmd *cobra.Command, args []string) {
			runNewProject()
			os.Exit(0)
		},
	}
	projectCmd.AddCommand(subNew)

	subInvoice := &cobra.Command{
		Use:   "invoice",
		Short: "setup invoice",
		Run: func(cmd *cobra.Command, args []string) {
			runNewProject()
			os.Exit(0)
		},
	}
	projectCmd.AddCommand(subInvoice)

	return projectCmd
}

func runNewProject() {

}
