package main

import (
  "github.com/freejaus/jauzi/cmd"
  "github.com/freejaus/jauzi/asciiart"
  "fmt"
)

func main() {
  asciiart.Header()

	cmd.Execute()

  asciiart.Hrule('=',65,true)
  fmt.Println("Quiting")
  asciiart.Hrule('=',65,true)
}
