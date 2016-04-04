package cmd

import (
//	"fmt"
	"github.com/spf13/cobra"
//  "github.com/freejaus/jauzi/lib"
)

// rmdbCmd represents the 'db rm' command
var rmdbCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command...`,
	//Run: func(cmd *cobra.Command, args []string) {
	//},
}

func init() {
	dbCmd.AddCommand(rmdbCmd)
}
