package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var name string
var age int
var nickname string
var rootCmd = cobra.Command{
	Use:   "dsw",
	Short: "this is for my test",
	Long:  `this is for my test use case`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("my name is %v,i am %d years old\n", name, age)
	},
}

var subCmd = cobra.Command{
	Use:   "nick name",
	Short: "this is for my test",
	Long:  `this is for my test use case`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("my nickname is %v,i am %d years old\n", nickname, age)
	},
}

func init() {

	rootCmd.Flags().StringVarP(&name, "name", "n", "luke", "get my name")
	rootCmd.Flags().IntVarP(&age, "age", "a", 10, "get my age")
	subCmd.Flags().StringVarP(&nickname, "nickname", "k", "luke", "get nick name")
	rootCmd.AddCommand(&subCmd)
}
func main() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
