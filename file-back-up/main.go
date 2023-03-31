package main

import (
	"file-back-up/filemd5"
	"fmt"
)

func main() {
	md5value, err := filemd5.FileMd5("./data.txt")
	if err != nil {
		fmt.Println("calc md5 error", err)
	} else {
		fmt.Printf("%x\n\n", md5value)
	}
}
