package readfile

import (
	"bufio"
	"flag"
	"io"
	"os"
)

func CheckInput() (io.ReadCloser, error) {
	fileInNameArg := flag.Arg(0)
	if len(fileInNameArg) > 0 {
		file, err := os.Open(fileInNameArg)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return os.Stdin, nil

}

func ChechOutput() (io.WriteCloser, error) {
	fileOutNameArg := flag.Arg(1)
	if len(fileOutNameArg) > 0 {
		file, err := os.OpenFile(flag.Arg(1), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return os.Stdout, nil
}

func ReadFile(file io.Reader) ([]string, error) {
	lines := []string{}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == io.EOF.Error() {
				return lines, nil
			}
			return nil, err
		}
		lines = append(lines, line)
	}
}

func WriteFile(file io.Writer, lines []string) (int, error) {
	writer := bufio.NewWriter(file)
	totalSize := 0
	for _, line := range lines {
		size, err := writer.WriteString(line)
		if err != nil {
			return 0, nil
		}
		totalSize += size
	}
	defer writer.Flush()
	return totalSize, nil
}
