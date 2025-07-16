/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fibbonacciCmd represents the fibbonacci command
var fibbonacciCmd = &cobra.Command{
	Use:   "fibbonacci",
	Short: "Generate a Fibonacci sequence up to a given number or count.",
	Long: `The Fibonacci command generates the Fibonacci sequence either up to a specified number or for a given count of terms.
	
	Example: matcli fibbonacci -l 5
	-l is the flag that gets the first 5 numbers from the fibbonacci series
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fibbonacci called")
	},
}

func GoFibbo(length int) ([]int , error){
	
}

func init() {
	rootCmd.AddCommand(fibbonacciCmd)
	fibbonacciCmd.Flags().IntP("length", "l", 7, "Gets the numbers from the fibbonacci series")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fibbonacciCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fibbonacciCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
