package arguments

import (
	"fmt"
)

type Arguments struct {
	Count               bool `flag:"c" usage: uniq -c`
	Duplicate           bool `flag:"d" usage: uniq -d`
	Uniq                bool `flag:"u" usage: uniq -u`
	Fields              int  `flag:"f" usage: uniq -f 4`
	Symbol              int  `flag:"s" usage: uniq -s 4`
	independentRegister bool `flag:"i" usage: uniq -i`
}

func (a Arguments) CheckArguments() {
	checkCount(a)
}

func checkCount(a Arguments) {
	if a.Count {
		fmt.Println("+")
	}
}
