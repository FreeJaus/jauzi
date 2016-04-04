package cmd

import (
//	"fmt"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command...`,
	//Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
	//	fmt.Println("db called")
	//},
}

var dbi,dbo,mfdroot,mfdfile string
var mfdrec bool

func init() {
	RootCmd.AddCommand(dbCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, and local flags which will only run when this command
	// is called directly:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
