package REPLs

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/REPLs/utils"
	"os"
)

func CBMShell() {
	for {
		fmt.Printf("%s > CBM :> ", os.Getenv("PWD"))
		_ = utils.GetLine()
	}
}
