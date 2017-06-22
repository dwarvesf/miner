package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func initEmailCommand() *cobra.Command {
	emailCmd := &cobra.Command{
		Use:   "email",
		Short: "send email with predefine template",
	}

	subHoliday := &cobra.Command{
		Use:   "holiday",
		Short: "send wishes to clients",
		Run: func(cmd *cobra.Command, args []string) {
			runEmailHoliday()
			os.Exit(0)
		},
	}
	emailCmd.AddCommand(subHoliday)

	subCandidate := &cobra.Command{
		Use:   "candidate",
		Short: "accept/reject candidate",
		Run: func(cmd *cobra.Command, args []string) {
			runEmailCandidate()
			os.Exit(0)
		},
	}
	emailCmd.AddCommand(subCandidate)

	subMeeting := &cobra.Command{
		Use:   "meeting",
		Short: "setup meeting",
		Run: func(cmd *cobra.Command, args []string) {
			runEmailMeeting()
			os.Exit(0)
		},
	}
	emailCmd.AddCommand(subMeeting)

	return emailCmd
}

func runEmailHoliday() {

}

func runEmailCandidate() {

}

func runEmailMeeting() {

}
