package main

import (
	"bytes"
	"fmt"
	"os"
)

const (
	// MaxBuffer maximum bytes of buffer for passing from file -> to file
	MaxBuffer int = 4096
)

func isEqual(fileName1, fileName2 string) bool {
	f1, err1 := os.Open(fileName1)
	if err1 != nil {
		fmt.Println("Read file 1 error:", err1)
	}
	defer f1.Close()

	f2, err2 := os.Open(fileName2)
	if err2 != nil {
		fmt.Println("Read file 2 error:", err2)
	}
	defer f2.Close()

	buf1 := make([]byte, MaxBuffer)
	buf2 := make([]byte, MaxBuffer)

	for {
		_, err1 = f1.Read(buf1)
		_, err2 = f2.Read(buf2)
		if err1 != nil || err2 != nil {
			break
		}
		if !bytes.Equal(buf1, buf2) {
			return false
		}
	}

	return true
}

func copy(fromFileName, toFileName string) {
	fromF, fromErr := os.Open(fromFileName)
	if fromErr != nil {
		fmt.Println("From file error:", fromErr)
	}
	defer fromF.Close()

	toF, toErr := os.Create(toFileName)
	if toErr != nil {
		fmt.Println("To file error:", toErr)
	}
	defer toF.Close()

	buf := make([]byte, MaxBuffer)

	for {
		n, err := fromF.Read(buf)
		if err != nil {
			break
		}

		// Not read full buffer
		actualBuffer := buf[:n]

		_, err = toF.Write(actualBuffer)
		if err != nil {
			break
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("./copy from_file to_file")
		return
	}

	copy(os.Args[1], os.Args[2])
}
