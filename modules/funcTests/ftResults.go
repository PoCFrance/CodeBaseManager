package funcTests

import (
    "fmt"
    "github.com/logrusorgru/aurora"
)

const (
    OK = iota
    NOK
    IGNORED

)

type ftResult struct {
    statusFail int
    stdoutFail int
    stderrFail int
}

func compareOutput(out, exp string) int {
    switch exp {
    case "":
        return IGNORED
    case out:
        return OK
    default:
        return NOK
    }
}

func (r *ftResult) CompareToRef(ref, my *ftExecution) {
    switch {
    case ref.status != my.status:
        r.statusFail = NOK
    case ref.outBuf.String() != my.outBuf.String():
        r.stdoutFail = NOK
    case ref.errBuf.String() != my.errBuf.String():
        r.stderrFail = NOK
    }
}

func (r *ftResult) CompareToExp(exp *ftExpected, my *ftExecution) {
    if exp.Status != my.status {
        r.statusFail = NOK
    }
    r.stdoutFail = compareOutput(exp.getExpOut(), my.outBuf.String())
    r.stderrFail = compareOutput(exp.getExpErr(), my.errBuf.String())
}

func (r *ftResult) isPerfect() bool {
    return r.statusFail == OK && r.stdoutFail == OK && r.stderrFail == OK
}

func (r *ftResult) showFailed() {
    //TODO: Improve error display (show difference)
    if r.statusFail == NOK {
        fmt.Println("status differs")
    }
    if r.stdoutFail == NOK {
        fmt.Println("stdout differs")

    }
    if r.stderrFail == NOK {
        fmt.Println("stderr differs")
    }
    fmt.Println("")
}

func (r *ftResult) Show(name string) {
    if r.isPerfect() {
        fmt.Printf("%s%s%s%s", aurora.Green("Success").Bold(), aurora.Bold(": ["), name,
            aurora.Bold("]\n"))
    } else {
        fmt.Printf("%s%s%s%s", aurora.Red("Failure").Bold(), aurora.Bold(": ["), name,
            aurora.Bold("]\n"))
        r.showFailed()
    }
}