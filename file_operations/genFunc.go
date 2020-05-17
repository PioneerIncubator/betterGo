package file_operations

import (
	"bufio"
	"bytes"
	"os"
)

func checkFileExists(filePath string) bool {
	var exist = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func WriteToFile(filePath, packageName string, input []byte) error {
	var f *os.File
	var err error
	var exist = true
	if checkFileExists(filePath) {
		f, err = os.OpenFile(filePath, os.O_APPEND, 0666)
	} else {
		exist = false
		f, err = os.Create(filePath)
	}
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
