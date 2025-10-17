package cmd

import (
	"fmt"
	"github.com/joho/godotenv"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	_ = godotenv.Load()
	fmt.Println("serve called")
}

func init() {
	rootCmd.AddCommand(serveCmd)

}
