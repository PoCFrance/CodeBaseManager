package build

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Helps you build your project regardless of language.",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("build called")
		},
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	parentCmd.AddCommand(buildCmd)
}
