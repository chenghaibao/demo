package root

import (
	"github.com/spf13/cobra"
)

func NewHttpGatewayCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "http",
		Short: "http gateway",
		Long:  "check and distribution",
		//RunE: func(cmd *cobra.Command, args []string) error {
		//	return runHttpGateway()
		//},
		Run: func(cmd *cobra.Command, args []string) {
			runHttpGateway()
		},
	}
	return cmd
}

func runHttpGateway() {

}
