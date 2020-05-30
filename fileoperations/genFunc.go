package fileoperations

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"regexp"
	"strings"
)

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckFuncExists(filePath string, listOfArgTypes []string) (bool, string) {
	if !checkFileExists(filePath) {
		return false, ""
	}

	for j, str := range listOfArgTypes {
		listOfArgTypes[j] = regexp.QuoteMeta(str)
	}

	// Generate `target`, which will be used to match the function name in `filepath`
	// It will be like `argname_1 int, argname_2 int\)`
	var target string
	switch length := len(listOfArgTypes); length {
	case 0:
		panic("Error:There is no argument in listOfArgTypes")
	case 1:
		// There is no comma behind %s because there just have only one arg
		target = fmt.Sprintf("argname_%d %s", 1, listOfArgTypes[0])
	default:
		// There is a comma behind %s because it's not the last arg
		target = fmt.Sprintf("argname_%d %s,", 1, listOfArgTypes[0])
		i := 1
		for ; i < length-2; i++ {
			// There is a comma behind %s because it's not the last arg
			target = fmt.Sprintf("%s argname_%d %s,", target, i+1, listOfArgTypes[i])
		}
		// There is no comma behind %s because it's the last arg
		target = fmt.Sprintf("%s argname_%d %s) %s", target, i+1, listOfArgTypes[i], listOfArgTypes[length-1])
	}
	target = regexp.QuoteMeta(target)

	fmt.Printf("Finding %s in %s...\n", target, filePath)
	funcExists, funcName := matchFunc(filePath, target)

	return funcExists, funcName
}

func matchFunc(filePath, origin string) (bool, string) {
	f, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return false, ""
			}
			panic(err)
		}

		if ok, _ := regexp.Match(origin, line); ok {
			fmt.Println("Function has been generated before!")
			funcName := getFuncNameFromLine(line)
			fmt.Println("Previous function name:", funcName)
			return true, funcName
		}
	}
}

func getFuncNameFromLine(line []byte) string {
	// line is like "func AddAB( argname_1 int, argname_2 int) int {"
	// then this func will match funcName which like "AddAB" in line
	expr := "func \\w+\\(" // regular expression
	reg, _ := regexp.Compile(expr)
	// matchRet is the result of regular expression match, it will like "func AddAB("
	matchRet := string(reg.Find(line))
	// funcName is like "AddAB"
	funcName := matchRet[5 : len(matchRet)-1]
	return funcName
}

func ensureDirExists(filePath string) error {
	s := strings.Split(filePath, "/")
	s = s[0 : len(s)-1]
	dirPath := strings.Join(s, "/")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func ensureFileExists(filePath string) (*os.File, bool, error) {
	var f *os.File
	var err error
	exist := false
	if err = ensureDirExists(filePath); err != nil {
		fmt.Println(err)
	}
	if checkFileExists(filePath) {
		exist = true
		f, err = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	} else {
		f, err = os.Create(filePath)
	}

	if err != nil {
		panic(err)
	}

	return f, exist, err
}

func WriteFuncToFile(filePath, packageName string, input []byte) error {
	var err error
	input, err = format.Source(input)
	if err != nil {
		fmt.Println(err)
	}
	f, exist, err := ensureFileExists(filePath)
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(f)
	if !exist {
		var buffer bytes.Buffer
		buffer.Write([]byte(packageName + "\n"))
		buffer.Write(input)
		input = buffer.Bytes()
	}
	if _, err = writer.Write(input); err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
