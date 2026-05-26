package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

//func main() {
//	fmt.Println("Hello file")
//}
