package file

import (
	"fmt"
	"os"
)

func Main() {
	/*
		FileInfo: 檔案訊息
		interface {
			Name(), 檔案名稱
			Size(), 檔案大小，文字長度
			Mode(), 權限
			八進制表示方式
			r ---> 004
			w ---> 002
			x ---> 001
			- ---> 000
			ModTime(), 修改時間
			IsDir(), 是否為目錄
			Sys()
		}
	*/
	file := "../practice/assets/docs/test.txt"
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println("err", err)
		// os.Stat: 檢查檔案是否存在
		// 下面是錯誤訊息
		// err stat ../practice/assets/docs/testa.txt: no such file or directory
		return
	}
	fmt.Printf("fileInfo is type %T\n", fileInfo)
	// 如果出現下面訊息，表示檔案有存在目錄底下
	// fileInfo is type *os.fileStat
	fmt.Println("fileInfo:")
	fmt.Println(fileInfo)
	// &{test.txt 2223 420 ....}
	fmt.Println("fileName:", fileInfo.Name())
	// fileName: test.txt
	fmt.Println("fileSize:", fileInfo.Size())
	// fileSize: 2223
	fmt.Println("fileMode:", fileInfo.Mode())
	// fileMode: -rw-r--r--
	fmt.Println("fileModTime:", fileInfo.ModTime())
	// fileModTime: 2023-03-07 10:10:58.2193228 +0800 CST
	// fileModTime: 2023-03-07 10:29:38.906598311 +0800 CST
	fmt.Println("fileIsDir:", fileInfo.IsDir())
	// fileIsDir: false
	fmt.Println("fileSys:", fileInfo.Sys())
	// fileSys: &{16777220 33188 1 129277281...}
	fmt.Println("filePerm:", fileInfo.Mode().Perm())
	// filePerm: -rw-r--r--
}

func HandleErr(err error) {
	if err != nil {
		fmt.Println("err", err)
		return
	}
}
