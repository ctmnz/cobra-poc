package stringer

import (
  "fmt"

  "example.com/stringer/pkg/stringer"
  "github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
  Use: "hello",
  Aliases: []string{"hi"},
  Short: "Just saying hi",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    res := stringer.SayHello(args[0])
    fmt.Println(res)
  },
}

func init() {
  rootCmd.AddCommand(helloCmd)
}

