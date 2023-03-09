package file

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// 檔案斷點續傳，讀或寫設置偏移量
func Seek(src string) {
	/*
		Seek(offset int64, whence int) (int64, error)，設置起始位子
		第一個參數：偏移量
		第二個參數：如何設置
	*/
	f, err := os.OpenFile(src, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 讀寫
	bs := []byte{0}
	f.Read(bs)
	fmt.Println(string(bs))

	f.Seek(4, io.SeekStart)
	f.Read(bs)
	fmt.Println(string(bs))

	f.Seek(2, 0) // SeekStart
	f.Read(bs)
	fmt.Println(string(bs))

	f.Seek(3, io.SeekCurrent)
	f.Read(bs)
	fmt.Println(string(bs))

	f.Seek(0, io.SeekEnd)
	f.WriteString("abs")

	// 從每次文件開啟時 io.SeekStart，插入的位子
	// 0 是指文件的最起始點，也就是讀寫時的起始位子
	// 如果有os.O_APPEND的參數，則會從文件末尾的位子開始讀寫
	// 如果只有 os.O_RDONLY 的參數，則會從起始位子開始讀寫，並覆蓋掉原先的資料同時插入新的資料
	f.Seek(0, io.SeekStart) // ABCDEFGHIJK
	f.WriteString("Apple")  // AppleFGHIJKabs

}

func SeekBreakPoint(src string) {
	/*
		斷點續傳:
		檔案傳遞：檔案複製
		複製到當前的位置：
		思路：邊複製，邊記錄複製總數量
	*/
	destFile := src[strings.LastIndex(src, "/")+1:]
	fmt.Println(destFile)
	// file.Seek(file2)
	tempFile := destFile + "temp.txt"
	fmt.Println(tempFile)

	origin, err := os.Open(src)
	HandleErr(err)
	dest, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandleErr(err)
	temp, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	HandleErr(err)

	defer origin.Close()
	defer dest.Close()

	// step1: 先讀取暫存檔中的資料，再seek到檔案的開頭
	temp.Seek(0, io.SeekStart)
	bs := make([]byte, 100)
	// 讀取檔案
	n1, _ := temp.Read(bs)
	// HandleErr(err) // err EOF 一開始的時候是空白的
	countStr := string(bs[:n1])
	count, _ := strconv.ParseInt(countStr, 10, 64)
	// HandleErr(err) // 無法正常解析countStr，如果一開始文件為空
	fmt.Println(count)

	// step2: 設置讀、寫的起始位子
	origin.Seek(count, io.SeekStart)
	dest.Seek(count, io.SeekStart)
	// 用來複製檔案內容
	data := make([]byte, 1024)
	// 讀取總數
	total := int(count)

	// step3: 複製檔案
	for {
		n2, err := origin.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("檔案複製完畢...", total)
			temp.Close()
			os.Remove(tempFile)
			break
		}
		n3, _ := dest.Write(data[:n2])
		total += n3

		// 將已複製的檔案內容，儲存到暫存檔案中
		temp.Seek(0, io.SeekStart)
		temp.WriteString(strconv.Itoa(total))

		fmt.Println("total:", total)

		// if total > 8000 {
		// 	panic("檔案傳輸中斷...")
		// }
	}

}
