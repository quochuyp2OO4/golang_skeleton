package cmd

import (
	"fmt"
	"os"

	"gocms/cmd/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocms",
	Short: "Gocms short",
	Long:  `Gocms is fast`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---> Root cms")
	},
}

func init() {
	rootCmd.AddCommand(server.ServerCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
