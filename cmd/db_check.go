package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
  "github.com/freejaus/jauzi/lib"
  "github.com/freejaus/jauzi/asciiart"
)

// checkdbCmd represents the 'db check' command
var checkdbCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command...`,
	Run: func(cmd *cobra.Command, args []string) {
		cs := lib.ReadDB(dbi,verb)

	  asciiart.Hrule('=',65,true)
	  fmt.Println("- Checking per byte equality:")
	  cmps, min := cs.CheckEq(verb)

	  fmt.Print(fmt.Sprintf("  MIN MATCH:\t%d\t%.2f\t",min[0],float32(min[0])*100/1024))
	  fmt.Println("(",(*cmps)[min[1]].A ,"|", (*cmps)[min[1]].B,")")

    if verb {
	   asciiart.Hrule('=',65,true)
	   fmt.Println("- Dumping diff:")
	   _,_ = cs.NotCommon(verb)
	  }
	},
}

func init() {
	dbCmd.AddCommand(checkdbCmd)
}
