package cmd

import (
	"baymax/config/config"
	"github.com/spf13/cobra"

	"baymax/server/grpc"
)

// grpcserverCmd represents the grpcserver command
var grpcserverCmd = &cobra.Command{
	Use:   "grpcserver",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		grpc.StartServer(cmd.Context(), config.Cfg.GRPCServerConfig)
	},
}

func init() {
	rootCmd.AddCommand(grpcserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
