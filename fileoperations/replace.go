package fileoperations

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Generate function calling statement by funName and arguments
func GenCallExpr(funName, assertType string, listOfArgs []string, isNew bool) string {
	callExpr := funName
	numOfArgs := len(listOfArgs)

	callExpr += "("
	for index, arg := range listOfArgs {
		callExpr += arg
		if index != numOfArgs-1 {
			callExpr += ", "
		}
	}

	callExpr += ")"
	if !isNew {
		if len(assertType) != 0 {
			callExpr += ".("
			callExpr += assertType
			callExpr += ")"
		}
	}

	if !isNew {
		callExpr = regexp.QuoteMeta(callExpr)
	}
	return callExpr
}

func ReplaceOriginFuncByFile(file, origin, target string) {
	output, needHandle, err := readFile(file, origin, target)
	if err != nil {
		panic(err)
	}
	if needHandle {
		err = writeCallExprToFile(file, output)
		if err != nil {
			panic(err)
		}
		fmt.Println(origin, "has been replaced with", target)

		// replace import statement
		dir, _ := os.Getwd()                                  // get current dir, equal to "pwd", like "/Users/.../src/.../test"
		gopath := fmt.Sprintf("%s/src/", os.Getenv("GOPATH")) // get env "GOPATH", like "/Users/.../src/"
		pkgName := strings.Split(origin, ".")[0]              // get package name, like "enum"
		oldPath := strings.Replace(dir, gopath, "", -1)       // oldPath == dir - gopath, like ".../test"
		oldImport := fmt.Sprintf("github.com/YongHaoWu/betterGo/%s", pkgName)
		newImport := fmt.Sprintf("%s/utils/%s", oldPath, pkgName)
		replaceOriginImport(file, oldImport, newImport)

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

func replaceOriginImport(file, origin, target string) {
	output, needHandle, err := readFile(file, origin, target)
	if err != nil {
		panic(err)
	}
	if needHandle {
		err = writeCallExprToFile(file, output)
		if err != nil {
			panic(err)
		}
		fmt.Println(origin, "has been replaced with", target)
	} else {
		fmt.Println("Can't find ", origin)
	}
}

// Read the file line by line to match origin and replace by target
func readFile(filePath, origin, target string) ([]byte, bool, error) {
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
			fmt.Println("Statement match success!")
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
}

// Write target function calling statement to the file
func writeCallExprToFile(filePath string, input []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
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
