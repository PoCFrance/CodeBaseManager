package funcTests

import (
	"github.com/PoCInnovation/CodeBaseManager/cli/modules/funcTests"
	"github.com/spf13/cobra"
)

func registerRun(parentCmd *cobra.Command) {
	var catCmd = &cobra.Command{
		Use:   "run [tests...]",
		Short: "Either run all tests or only the specified ones.",
		Run: func(_ *cobra.Command, args []string) {
			funcTests.Run(args)
		},
	}

	parentCmd.AddCommand(catCmd)
}
