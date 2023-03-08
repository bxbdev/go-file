package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CopyFile(src, dst string) (int, error) {
	origin, err := os.Open(src)
	if err != nil {
		return 0, err
	}

	dest, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer origin.Close()
	defer dest.Close()

	// write
	bs := make([]byte, 1024) // No need make([]byte, 1024, 1024), Let Go language determine the size and capacity of the slice automatically.
	// n := -1 // remove this line
	total := 0
	for {
		n, err := origin.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("檔案複製完成")
			break
		} else if err != nil {
			fmt.Println("檔案讀取錯誤")
			return total, err
		}
		total += n
		dest.Write(bs[:n])
	}

	return total, nil
}

func CpyFile(src, dst string) (int64, error) {
	origin, err := os.Open(src)
	if err != nil {
		return 0, err
	}

	dest, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer origin.Close()
	defer dest.Close()

	return io.Copy(dest, origin)
}

// ioutil的讀取跟寫入都是一次性的，檔案過大的話會溢出
// 建議使用 CopyFile, CpyFile這二種比較安全的作法
func IoutilCpyFile(src, dst string) (int, error) {
	bs, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, err
	}

	writeErr := ioutil.WriteFile(dst, bs, os.ModePerm)
	if writeErr != nil {
		return 0, writeErr
	}
	return len(bs), nil
}

// ChatGPT recommend, refactor for IoutilCpyFile func
func CpFile(src, dst string) (int64, error) {
	sourceFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destFile.Close()

	bytesCopied, err := io.Copy(destFile, sourceFile)
	if err != nil {
		return 0, err
	}
	return bytesCopied, nil
}
