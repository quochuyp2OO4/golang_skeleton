package cmd

import (
	"fmt"
	"gocms/cmd/server"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Gocms",
	Short: "Gocms fast",
	Long:  `Root gocms run`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root cms")
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
