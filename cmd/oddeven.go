/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// oddevenCmd represents the oddeven command
var oddevenCmd = &cobra.Command{
	Use:   "oddeven",
	Short: "Check if a number is odd or even.",
	Long: `The odd-even command checks whether the given number is odd or even.

Example:
  mathcli oddeven -n 42
  Output: 42 is even

This command uses simple modulo logic to determine the parity of the number — a handy utility for quick math checks.`,
	Run: func(cmd *cobra.Command, args []string) {
		number,_:=cmd.Flags().GetInt("parity")
		res, err:= GoParity(number)
		if err!=nil {
			fmt.Println(err)
		}else {
			fmt.Println(res)
		}
	},
}

func GoParity(num int) (string, error){
	if num == 0 || num<0 {
		return "", errors.New("no negative number")
	}

	if num % 2 == 0{
		return "The Number is Even", nil
	}else {
		return "The Number is Odd", nil
	}
}

func init() {
	rootCmd.AddCommand(oddevenCmd)

	oddevenCmd.Flags().IntP("parity", "p", 0, "checks if the number is odd or even")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oddevenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oddevenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
