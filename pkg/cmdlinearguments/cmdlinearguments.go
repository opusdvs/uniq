package cmdlinearguments

import (
	"fmt"
)

type CmdLineArguments struct {
	Count               bool `flag:"c" usage: uniq -c`
	Duplicate           bool `flag:"d" usage: uniq -d`
	Uniq                bool `flag:"u" usage: uniq -u`
	Fields              int  `flag:"f" usage: uniq -f 4`
	Symbol              int  `flag:"s" usage: uniq -s 4`
	independentRegister bool `flag:"i" usage: uniq -i`
}

func (a CmdLineArguments) CheckArguments() {
	checkCount(a)
}

func checkCount(a CmdLineArguments) {
	if a.Count {
		fmt.Println("+")
	}
}
