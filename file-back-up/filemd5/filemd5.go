package filemd5

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
)

func FileMd5(filepath string) ([]byte, error) {
	md5_inst := md5.New()
	fp, err := os.Open(filepath)
	if err != nil {
		fmt.Println("open file error")
		return nil, errors.New("opon file error")
	}
	data_buf := make([]byte, 4096)
	reader := bufio.NewReader(fp)
	for {
		_, err := reader.Read(data_buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read file error")
				return nil, errors.New("read file error")
			} else {
				break
			}
		}
		md5_inst.Write(data_buf)
	}
	return md5_inst.Sum(nil), nil
}
