package config

import (
	"github.com/spf13/cobra"
	"nacosctl/printer"
	"nacosctl/process/config"
	"nacosctl/process/constant"
)

var GetConfigCmd = &cobra.Command{
	Use:     "config [dataId]",
	Short:   "Get configuration information",
	Long:    "Get configuration information",
	Example: "nacosctl get config [dataId]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			printer.Red("[error]:dataId must not be null")
			printer.Yellow("[see]:nacosctl get config -h")
			return
		}
		config.ParseGetConfigCmd(cmd, args[0], configNamespaceId, group)
	},
}

func init() {
	GetConfigCmd.Flags().StringVarP(&configNamespaceId, "namespaceId", "i", "", "namespace id")
	GetConfigCmd.Flags().StringVarP(&group, "group", "g", constant.DefaultGroup, "config group")
}
