package file_operations

import (
	"bufio"
	"bytes"
	"os"
)

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func ensureFileExists(filePath string) (*os.File, error, bool) {
	var f *os.File
	var err error
	var exist = false
	if checkFileExists(filePath) {
		exist = true
		f, err = os.OpenFile(filePath, os.O_APPEND, 0666)
	} else {
		f, err = os.Create(filePath)
	}

	if err != nil {
		panic(err)
		return nil, err, exist
	}

	return f, err, exist
}

func WriteFuncToFile(filePath, packageName string, input []byte) error {
	f, err, exist := ensureFileExists(filePath)
	defer f.Close()
	if err != nil {
		panic(err)
		return err
	}

	writer := bufio.NewWriter(f)
	if !exist {
		tmpStr := packageName + "\n"
		tmpBuffer := []byte(tmpStr)
		var buffer bytes.Buffer
		buffer.Write(tmpBuffer)
		buffer.Write(input)
		input = buffer.Bytes()
	}
	_, err = writer.Write(input)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
