/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var fileName string

// touchCmd represents the touch command
var touchCmd = &cobra.Command{
	Use:   "touch",
	Short: "Create new file",
	Long: `Create a new file using the touch command. The default file name is text.txt.
	Ex: tool touch <filename>`,
	Run:               CreateFile,
	Version:           Version,
	Args:              cobra.MaximumNArgs(1),
	DisableAutoGenTag: true,
}

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Read file",
	Long:  `Read a file by using the cat command.Ex: tool cat <filename>`,
	Run:   ReadFile,
}

func CreateFile(cmd *cobra.Command, args []string) {
	var fileName = "text.txt"

	if len(args) > 0 && args[0] != "" {
		fileName = args[0]
	}
	// check if the file already exists
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("Error: %s already exists", fileName)
		return
	}
	os.Create(fileName)
}

func ReadFile(cmd *cobra.Command, args []string) {
	if len(args) > 0 && args[0] != "" {
		fileName = args[0]
	}
	// check if there is a file
	if _, err := os.Stat(fileName); err == nil {
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(fileData))
	} else {
		fmt.Println("Error:", err)
	}
}

func init() {
	rootCmd.AddCommand(touchCmd)
	rootCmd.AddCommand(catCmd)

	// catCmd.SetUsageFunc(func(c *cobra.Command) error {
	// 	fmt.Println(" tool touch [flags]<argument>")
	// 	return nil
	// })
	// catCmd.SetUsageTemplate(" tool touch <argument>")
	catCmd.Flags().StringVarP(&fileName, "file", "f", "", "Read a file")
}
