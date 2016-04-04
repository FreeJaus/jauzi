package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/freejaus/jauzi/lib"
)

// appenddbCmd represents the 'db append' command
var appenddbCmd = &cobra.Command{
	Use:   "append",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command...`,
	Run: func(cmd *cobra.Command, args []string) {
	 if 0==len(adbi) { fmt.Println("ERROR: Input file not provided."); return; }
	 cs := lib.ReadDB(adbi,verb)
   var w bool = false
   if len(amfdroot)!=0 {
		w = true
		n := len(cs)
	  cs.AppendWalk(amfdroot, amfdrec, verb)
	  fmt.Println("...",len(cs)-n,"cards added")
   }
	 if len(amfdfile)!=0 {
    w = true
		fmt.Printf("- Appending card in '%s' to the DB...\n",amfdfile)
		cs.Append(amfdfile,verb)
	 }
   if w { cs.WriteDB(adbo) } else { fmt.Println("WARNING: No data to write") }
	},
}

var adbi,adbo,amfdroot,amfdfile string
var amfdrec bool

func init() {
	dbCmd.AddCommand(appenddbCmd)

  appenddbCmd.Flags().StringVarP(&adbi, "input", "i", "", "input file, to append data to")
	appenddbCmd.Flags().StringVarP(&amfdroot, "dir", "d", "", "search root directory")
	appenddbCmd.Flags().BoolVarP(&amfdrec, "recursive", "r", false, "recursive search")
	appenddbCmd.Flags().StringVarP(&amfdfile, "card", "c", "", "card file to append to DB")
	appenddbCmd.Flags().StringVarP(&adbo, "output", "o", "appended.cdb", "output file")
}
