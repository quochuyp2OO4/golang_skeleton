package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "Server",
	Short: "Server gocms fast",
	Long:  `Server gocms run`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server cms")
	},
}
