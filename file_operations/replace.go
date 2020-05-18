package file_operations

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

// Generate function calling statement by funName and arguments
func GenerateCall(funName string, args []string, isNew bool, assertType string) string {
	lenOfArgs := len(args)
	if isNew {
		funName += "("
	} else {
		funName += "\\("
	}
	for index, arg := range args {
		funName += arg
		if index != lenOfArgs-1 {
			funName += ", "
		}
	}
	if isNew {
		funName += ")"
	} else {
		funName += "\\)"
		lenOfAssert := len(assertType)
		if lenOfAssert != 0 {
			funName += "\\.\\("
			funName += assertType
			funName += "\\)"
		}
	}
	return funName
}

func ReplaceOriginFuncByFile(file, origin, target string) {
	output, needHandle, err := readFile(file, origin, target)
	if err != nil {
		panic(err)
	}
	if needHandle {
		err = writeCallToFile(file, output)
		if err != nil {
			panic(err)
		}
		fmt.Println(origin, "has been replaced with", target)
	} else {
		fmt.Println("Can't find ", origin)
	}
}

func ReplaceOriginFuncByDir(path, origin, target string) {
	files := getFiles(path)
	for _, file := range files {
		fmt.Println("File:", file, "is been replacing...")
		ReplaceOriginFuncByFile(file, origin, target)
		fmt.Println("File:", file, "...done")
	}
}

// Read the file line by line to match origin and replace by target
func readFile(filePath string, origin string, target string) ([]byte, bool, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, false, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	needHandle := false
	output := make([]byte, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, needHandle, nil
			}
			return nil, needHandle, err
		}

		if ok, _ := regexp.Match(origin, line); ok {
			fmt.Println("Match success!")
			reg := regexp.MustCompile(origin)
			newByte := reg.ReplaceAll(line, []byte(target))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
			if !needHandle {
				needHandle = true
			}
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
	return output, needHandle, nil
}

// Write target function calling statement to the file
func writeCallToFile(filePath string, input []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(input)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func getFiles(path string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return files
}
