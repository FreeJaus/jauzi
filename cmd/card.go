package cmd

import (
//	"fmt"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command...`,
	//Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
	//	fmt.Println("card called")
	//},
}

func init() {
	RootCmd.AddCommand(cardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
