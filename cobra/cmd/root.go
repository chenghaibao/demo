package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"main.go/cmd/root"
	"os"
)

var RootCmd *cobra.Command

const Version = "1.0.0"

var Build = "local_build"

func init() {
	RootCmd = &cobra.Command{
		// 基础命令
		Use:              "chartGateway",
		Short:            "shor description",
		Long:             "description",
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
		Args:             noArgs, //验证器
		RunE: func(cmd *cobra.Command, args []string) error {
			return ShowHelp(os.Stderr)(cmd, args)
		},
		Version:               fmt.Sprintf("%s, build %s", Version, Build),
		DisableFlagsInUseLine: true,
	}
	RootCmd.AddCommand(root.NewHttpGatewayCommand())
}

//ShowHelp 查看命令行帮助.
func ShowHelp(err io.Writer) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetErr(err)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}

//命令行没有网关提示错误
func noArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return nil
	}
	return fmt.Errorf("chart-gateway: '%s' is not a gateway command.\nSee 'chart-gateway --help'", args[0])
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
