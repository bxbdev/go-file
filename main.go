package main

import "go-file/file"

// "os"

func main() {
	// 檔案資訊
	// file.Main()

	// 檔案路徑
	// file.FilePath()

	// 製作目錄
	// file.MakeFolder()

	// 新建檔案
	// 需要建立的檔案路徑+檔名.副檔名
	// file1 := "../go-file/assets/docs/test_102a.txt"
	// file2 := "test.txt"

	// 檢查檔案是否存
	// checkFile, _ := os.Stat(file1)
	// 如果checkFile已有取得資料就判定檔案已經存在
	// if checkFile != nil {
	// 	fmt.Println("檔案已存在")
	// 	return
	// }
	// 建立新的檔案
	// file.CreateFile(file1)
	// file.CreateFile(file2)

	// 變更檔名
	// newFileName := "../go-file/assets/docs/test1.txt"
	// file.Rename(file1, newFileName)

	// 刪除檔案
	// file.DeleteFile(file1)

	// 刪除目錄
	// file.DeleteFolder("../go-file/assets")

	// 刪除目錄包含以下檔案
	// 此指令需要先檢查該目錄是否存在，不然都會直接刪除成功
	// 需要注意此指令是完全刪除檔案和目錄，需要特別小心
	// file.ForceDelete("../go-file/assets")

	// story2 := "../go-file/file/story2.txt"
	// test := "../go-file/assets/docs/test.txt"
	// ab := "../go-file/assets/docs/ab.txt"
	// 開啟檔案 ReadOnly
	// file.OpenFile(story2)
	// file.ReadFile(story2)

	// 開啟檔案 ReadWrite
	// file.OpenFileWrite(file1)
	// file.WriteFile(ab, &[]byte{}, "Hello World")

	// 複製文件
	// srcFile := "../go-file/file/story.txt"
	// destFile := "story_gpt.txt"

	// total, _ := file.CopyFile(srcFile, destFile)
	// fmt.Println(total)

	// 複製圖片
	// srcImg := "../go-file/assets/images/andy-holmes-LUpDjlJv4_c-unsplash.jpg"
	// destImg := "star.jpg"

	// cp, err := file.CpyFile(srcImg, destImg)
	// cp, err := file.IoutilCpyFile(srcImg, destImg)
	// cp, err := file.CpFile(srcImg, destImg)
	// file.HandleErr(err)
	// fmt.Println(cp)
	// fmt.Println("檔案複製完成")

	// 使用緩衝功能將內容複製到目標檔案
	// file.Bufio(ab, "Buy a milk")

	// ioutil讀取檔案
	// file.IoutilRead(test)

	// ioutil寫入檔案
	// str := "風花雪月，雲淡風清"
	// cc := "../go-file/assets/docs/cc.txt"
	// file.IoutilWrite(cc, str)

	// ioutil取讀所有內容
	// str := "New research shows that coffee may have health benefits. Drinking moderate amounts of coffee could reduce the risk of heart disease."
	// file.IoutilReadAll(str)

	// ioutil讀取資料夾
	folder := "../go-file"
	file.IoutilReadDir(folder)
}
