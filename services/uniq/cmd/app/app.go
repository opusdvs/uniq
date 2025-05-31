package main

import (
	"fmt"

	"github.com/Benzinga/flagstruct"
	"github.com/opusdvs/uniq/pkg/cmdlinearguments"
	"github.com/opusdvs/uniq/pkg/readfile"
	"github.com/opusdvs/uniq/pkg/uniqstrings"
)

func main() {
	var arg cmdlinearguments.CmdLineArguments

	err := flagstruct.Configure(&arg)
	if err != nil {
		err = fmt.Errorf("error configure cmd line arguments: %v", err.Error())
		fmt.Println(err.Error())
		return
	}

	inFile, err := readfile.CheckInput()
	if err != nil {
		err = fmt.Errorf("error input file: %v", err.Error())
		fmt.Print(err.Error())
		return
	}
	defer inFile.Close()

	outFile, err := readfile.ChechOutput()
	if err != nil {
		err = fmt.Errorf("error output file: %v", err.Error())
		fmt.Print(err.Error())
		return
	}
	defer outFile.Close()

	lines, err := readfile.ReadFile(inFile)
	if err != nil {
		err = fmt.Errorf("error read file: %v", err.Error())
		fmt.Print(err.Error())
		return
	}

	uniqLines := uniqstrings.UniqString(lines, arg)
	_, err = readfile.WriteFile(outFile, uniqLines)
	if err != nil {
		err = fmt.Errorf("error write file: %v", err.Error())
		fmt.Print(err.Error())
		return
	}
}
