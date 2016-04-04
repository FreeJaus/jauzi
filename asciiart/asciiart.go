package asciiart

import (
  "fmt"
)

//Header prints the header to the terminal.
func Header(){
  Hrule('=',65,true)
  Hrule(' ',19,false); fmt.Println("    | | __ _ _   _ ___(_)");
  Hrule(' ',19,false); fmt.Println(" _  | |/ _` | | | |_  / |");
  Hrule(' ',19,false); fmt.Println("| |_| | (_| | |_| |/ /| |");
  Hrule(' ',19,false); fmt.Println(" \\___/ \\__,_|\\__,_/___|_|");
  Hrule('=',65,true)
}

// Hurel prints m times the byte b and, if ln is true, an EOL is printed
// at the end.
func Hrule(b byte, m int, ln bool) {
  for i := 0; i < m; i++ { fmt.Print(string(b)) };
  if ln { fmt.Println("") }
}
