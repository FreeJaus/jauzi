package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/freejaus/jauzi/lib"
)

// newdbCmd represents the 'db new' command
var newdbCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command...`,
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("- Creating new card DB")
	  cs := lib.NewCardDB()
    cs.AppendWalk(mfdroot, mfdrec, verb)
		fmt.Println("...",len(cs),"cards added")
    cs.WriteDB(dbo)
	},
}

func init() {
	dbCmd.AddCommand(newdbCmd)

  newdbCmd.Flags().StringVarP(&mfdroot, "dir", "d", "./", "search root directory")
	newdbCmd.Flags().BoolVarP(&mfdrec, "recursive", "r", false, "recursive search")
	newdbCmd.Flags().StringVarP(&dbo, "output", "o", "new.cdb", "output file")
}
