package main

import (
	"fmt"
	"go-file/file"
	"os"
)

func main() {
	// 檔案資訊
	file.Main()

	// 檔案路徑
	file.FilePath()

	// 製作目錄
	file.MakeFolder()

	// 新建檔案
	// 需要建立的檔案路徑+檔名.副檔名
	file1 := "../go-file/assets/docs/test_102a.txt"
	file2 := "test.txt"
	// 檢查檔案是否存
	checkFile, _ := os.Stat(file1)
	// 如果checkFile已有取得資料就判定檔案已經存在
	if checkFile != nil {
		fmt.Println("檔案已存在")
		return
	}
	// 建立新的檔案
	file.CreateFile(file1)
	file.CreateFile(file2)

	// 開啟檔案 ReadOnly
	file.OpenFile(file1)

	// 開啟檔案 ReadWrite
	file.OpenFileWrite(file1)

	// 變更檔名
	newFileName := "../go-file/assets/docs/test1.txt"
	file.Rename(file1, newFileName)

	// 刪除檔案
	file.DeleteFile(file1)

	// 刪除目錄
	file.DeleteFolder("../go-file/assets")

	// 刪除目錄包含以下檔案
	// 此指令需要先檢查該目錄是否存在，不然都會直接刪除成功
	// 需要注意此指令是完全刪除檔案和目錄，需要特別小心
	file.ForceDelete("../go-file/assets")

}
