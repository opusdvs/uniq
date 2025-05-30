package main

import (
	"fmt"

	"github.com/Benzinga/flagstruct"
	"github.com/opusdvs/uniq/pkg/arguments"
	"github.com/opusdvs/uniq/pkg/readfile"
	"github.com/opusdvs/uniq/pkg/uniqstrings"
)

func main() {
	var arg arguments.Arguments

	err := flagstruct.Configure(&arg)
	if err != nil {
		fmt.Println(err)
		return
	}

	inFile, err := readfile.CheckInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inFile.Close()

	outFile, err := readfile.ChechOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	lines, err := readfile.ReadFile(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	uniqLines := uniqstrings.UniqString(lines, arg)
	_, err = readfile.WriteFile(outFile, uniqLines)
	if err != nil {
		fmt.Println(err)
		return
	}

}
