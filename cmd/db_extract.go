package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/freejaus/jauzi/lib"
)

// extractdbCmd represents the 'db extract' command
var extractdbCmd = &cobra.Command{
	Use:   "extract",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command...`,
	Run: func(cmd *cobra.Command, args []string) {
   if 0==len(edbi) { fmt.Println("ERROR: Input file not provided."); return; }
   if 0==len(eclist) { fmt.Println("ERROR: Card list not provided."); return; }
   cs := lib.ReadDB(edbi,verb)
   ids := cs.ParseList(&eclist)
   switch len(ids) {
    case 0: fmt.Println("WARNING: No data to write")
		case 1:
		 cs[ids[0]].Write(edbo+".mfd", verb)
		default:
		 ecs := cs.Extract(ids)
		 ecs.WriteDB(edbo+".cdb")
	 }
	},
}

var edbi,edbo,eclist string

func init() {
	dbCmd.AddCommand(extractdbCmd)

	extractdbCmd.Flags().StringVarP(&edbi, "input", "i", "", "input DB file, to extract data from")
	extractdbCmd.Flags().StringVarP(&eclist, "list", "l", "", "comma separated list of names or id numbers of the cards to be extracted")
	extractdbCmd.Flags().StringVarP(&edbo, "output", "o", "extracted", "output filename (without extension, which will be either '.cdb' or '.mfd')")
}
