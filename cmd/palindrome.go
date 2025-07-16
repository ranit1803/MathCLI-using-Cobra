/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// palindromeCmd represents the palindrome command
var palindromeCmd = &cobra.Command{
	Use:   "palindrome",
	Short: "Check if a given number is a palindrome.",
	Long: `The palindrome command checks whether a given positive number reads the same forward and backward.

Example:
  mathcli palindrome -n 121
  Output: Palindrome true

A number is a palindrome if it remains the same when its digits are reversed. This command helps you test such numbers quickly from the CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		num, _:= cmd.Flags().GetInt("number")
		palin, err:= GoPalin(num)
		if err!=nil{
			fmt.Println(err)
		}else {
			fmt.Println(palin)
		}
	},
}

func GoPalin(num int) (bool, error){
	if num< 0{
		return false, errors.New("no negative numbers")
	}
	rev, err:= GoReverse(num)
	if err != nil{
		fmt.Println("Not a Palindrome")
		return false, err
	}
	fmt.Println("Palindrome")
	return rev == num, nil
}

func init() {
	rootCmd.AddCommand(palindromeCmd)

	palindromeCmd.Flags().IntP("number", "n", 0, "checks number is palindrome")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// palindromeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// palindromeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
