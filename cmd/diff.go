package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command...`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("diff called")

		//cards := db.ReadDB(dbi)
		//paircmp := db.NewCardDB()
	  //paircmp = append(paircmp,cards[4],cards[8])
	  //_,_ = db.NotCommon(&paircmp,true)

	},
}

func init() {
	RootCmd.AddCommand(diffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
