/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// reverseCmd represents the reverse command
var reverseCmd = &cobra.Command{
	Use:   "reverse",
	Short: "Reverse the digits of a given number.",
	Long: `The reverse command takes a positive integer and returns its digits in reverse order.

Example:
  mathcli reverse -n 1234
  Output: 4321

This command is part of the Math CLI project built to practice common math-based coding problems using Cobra in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		num, _:= cmd.Flags().GetInt("number")
		rev, err:= GoReverse(num)
		if err!= nil{
			fmt.Println(err)
		}else {
			fmt.Println("The original number is:",num)
			fmt.Println("The reversed number is:",rev)
		}
	},
}

func GoReverse(num int) (int, error){
	if num<0 {
		return 0, errors.New("only positive number allowed")
	}
	rev:= 0
	for num> 0{
		last_digit:= num % 10
		rev = (rev *10) + last_digit
		num/=10
	}
	return rev,nil
}

func init() {
	rootCmd.AddCommand(reverseCmd)

	reverseCmd.Flags().IntP("number", "n", 0, "reverses a number")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reverseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reverseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
