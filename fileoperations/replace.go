package fileoperations

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
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
			callExpr = fmt.Sprintf("%s.(%s)", callExpr, assertType)
		}
		callExpr = regexp.QuoteMeta(callExpr)
	}

	return callExpr
}

func ReplaceOriginFuncByFile(file, origin, target string) {
	output, needHandle, err := readFile(file, origin, target)
	if err != nil {
		log.Fatal(err)
	}
	if needHandle {
		err = writeCallExprToFile(file, output)
		if err != nil {
			log.Fatal(err)
		}
		log.WithFields(log.Fields{
			"originCallExpr": origin,
			"targetCallExpr": target,
		}).Info("Replace function call expression successfully")

		// replace import statement
		dir, _ := os.Getwd()                                  // get current dir, equal to "pwd", like "/Users/.../src/.../test"
		gopath := fmt.Sprintf("%s/src/", os.Getenv("GOPATH")) // get env "GOPATH", like "/Users/.../src/"
		pkgName := strings.Split(origin, ".")[0]              // get package name, like "enum"
		pkgName = strings.TrimRight(pkgName, "\\")
		oldPath := strings.ReplaceAll(dir, gopath, "") // oldPath == dir - gopath, like ".../test"
		oldImport := fmt.Sprintf("\"github.com/PioneerIncubator/betterGo/%s\"", pkgName)
		newImport := fmt.Sprintf("%s \"%s/utils/%s\"", pkgName, oldPath, pkgName)
		replaceOriginImport(file, oldImport, newImport)

	} else {
		log.WithFields(log.Fields{
			"originCallExpr": origin,
		}).Error("Replace function call expression failed, the expr to be replaced was not found!")
	}
}

func ReplaceOriginFuncByDir(path, origin, target string) {
	files := getFiles(path)
	for _, file := range files {
		log.WithFields(log.Fields{
			"fileName": file,
		}).Info("The function call expression is being replaced")
		ReplaceOriginFuncByFile(file, origin, target)
	}
}

func replaceOriginImport(file, origin, target string) {
	origin = regexp.QuoteMeta(origin)
	output, needHandle, err := readFile(file, origin, target)
	if err != nil {
		log.Fatal(err)
	}
	if needHandle {
		err = writeCallExprToFile(file, output)
		if err != nil {
			log.Fatal(err)
		}
		log.WithFields(log.Fields{
			"originImportStmt": origin,
			"targetImportStmt": target,
		}).Info("Replace import statement successfully")
	} else {
		log.WithFields(log.Fields{
			"originImportStmt": origin,
		}).Error("Replace import statement failed, the stmt to be replaced was not found!")
	}
}

// Read the file line by line to match origin and replace by target
func readFile(filePath, origin, target string) ([]byte, bool, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, false, err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Error(err)
		}
	}()
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
			log.WithFields(log.Fields{
				"filePath":  filePath,
				"statement": origin,
			}).Info("The statement was found from the file")
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
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Error(err)
		}
	}()
	writer := bufio.NewWriter(f)
	_, err = writer.Write(input)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
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
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Error getting file from directory")
	}
	return files
}
