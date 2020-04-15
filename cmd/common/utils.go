package common

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func IsCBMRepository(_ *cobra.Command, av []string) error {
	const (
		currentDir = 0
		goToDir    = 1
	)

	switch len(av) {
	case currentDir:
		st, err := os.Stat(".cbm/")
		if err != nil || !st.IsDir() {
			return errors.New("Not in a CBM Repository.")
		}
	case goToDir:
		st, err := os.Stat(av[0])
		if err != nil || !st.IsDir() {
			return errors.New("Invalid filepath")
		}
		st, err = os.Stat(av[0] + "/.cbm/")
		if err != nil || !st.IsDir() {
			return errors.New("Not a CBM Repository")
		}
		if err = os.Chdir(av[0]); err != nil {
			return errors.Unwrap(err)
		}
	}
	return nil
}

func RetrieveSubCommandsNames(cmd *cobra.Command) []string {
	subCmdsName := []string{}
	subCmds := cmd.Commands()
	for _, c := range subCmds {
		subCmdsName = append(subCmdsName, strings.Fields(c.Use)[0])
	}
	return subCmdsName
}