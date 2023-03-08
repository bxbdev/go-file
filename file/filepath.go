package file

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func FilePath() {
	fileName1 := "/xxx/xxx/xxx/practice/assets/docs/test.txt"
	fileName2 := "story2.txt"
	// 檢測是否為絕對路徑
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))
	// 檢查目前的檔案絕對路徑
	fmt.Println(filepath.Abs(fileName1))
	// /xxx/xxx/xxx/practice/assets/docs/test.txt <nil>
	fmt.Println(filepath.Abs(fileName2))
	// /xxx/xxx/xxx/go/practice/story2.txt <nil>

	// 取得父目錄路徑 兩種寫法
	fmt.Println("取得父資料夾路徑:", filepath.Dir(fileName1))
	fmt.Println("取得父資料夾路徑:", path.Join(fileName1, ".."))
}

// 創建資料夾
func MakeFolder() {
	// 只能建立單層資料夾
	// err := os.Mkdir("../practice/test", os.ModePerm)
	// if err != nil {
	// 	fmt.Println("資料夾建立失敗，資料夾已經存在")
	// 	fmt.Println("err", err)
	// 	return
	// }
	// fmt.Println("資料夾建立成功")

	// 當資料夾不存在時，直接連同不存在的資料夾一起新建
	err := os.MkdirAll("../practice/test1/testA", os.ModePerm)
	if err != nil {
		fmt.Println("資料夾建立失敗，資料夾已經存在(test1) 或 資料夾不存在(test1)")
		fmt.Println("err", err)
		return
	}
	fmt.Println("多層資料夾建立成功")
}

// 新建檔案
func CreateFile(f string) {
	file, err := os.Create(f)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("檔案建立成功")
	// path.Base() 取得最後一個路徑名稱，如果沒有接檔案，只有資料夾名稱，
	// 那麼只會顯示最後一個資料夾的名稱
	// ex: /xxx/xxx/xxx/practice/assets/docs/test.txt
	// 這裡會取得的是test.txt
	// ex: /xxx/xxx/xxx/practice/assets/docs/
	// 這裡會取得的是docs
	fmt.Println("檔案:", path.Base(f), file)
}

// 打開檔案
func OpenFile(f string) {
	// 只能讀取檔案
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("打開檔案失敗")
		fmt.Println("err", err)
		return
	}
	fmt.Println("打開檔案成功", file)
	fmt.Println("讀取檔案開始")
	defer file.Close()
}

// 打開的檔案可讀取和寫入
func OpenFileWrite(f string) {
	file, err := os.OpenFile(f, os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("打開檔案失敗")
		fmt.Println("err", err)
		return
	}
	fmt.Println("打開檔案成功", file)
	fmt.Println("編輯檔案開始")
	defer file.Close()
}

/*
	os.OpenFile(檔案名稱, 檔案模式, 檔案權限)
	第一參數為檔案名稱
	第二參數為檔案開啟模式
	const (
		Exactly one of RDONLY, O_WRONLY, or O_RDWR must be specified.
		O_RDONLY int = syscall.O_RDONLY // open the file read-only
		O_WRONLY int = syscall.O_WRONLY // open the file write-only
		O_RDWR   int = syscall.O_RDWR   // open the file read-write

		The remaining values may be or ed in to control behavior.
		O_APPEND int = syscall.O_APPEND // append data to the file when writing
		O_CREATE int = syscall.O_CREAT  // create the file if it does not exist
		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O
		O_TRUNC  int = syscall.O_TRUNC  // truncate regular writeable file when opened
	)
	第三檔案權限
	os.ModePerm (ex:  0777, 0775, 0666, 0644, 0555 etc.. )
	可以直接寫成數字型態

	關閉檔案與程式之間的繫結
	defer file.Close()
*/

func DeleteFile(f string) {
	err := os.Remove(f)
	if err != nil {
		fmt.Println("刪除檔案失敗")
		fmt.Println("err", err)
		return
	}
	fmt.Println("刪除檔案成功", path.Base(f))
}

func DeleteFolder(f string) {
	// 只能刪除空目錄，如果底下有檔案，需要先刪除檔案
	err := os.Remove(f)
	if err != nil {
		fmt.Println("刪除目錄失敗")
		fmt.Println("err", err)
		return
	}
	fmt.Println("刪除目錄成功", path.Base(f))
}

// 直接刪除目錄和檔案，不會直接進到垃圾桶
func ForceDelete(path string) {
	_, err_c := os.Stat(path)
	if os.IsNotExist(err_c) {
		fmt.Println("目錄不存在", err_c)
		return
	}

	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("刪除目錄失敗")
		fmt.Println("err", err)
		return
	}
	fmt.Println("刪除目錄成功", path)
}

func Rename(o string, n string) {

	_, err_o := os.Stat(o)
	if os.IsNotExist(err_o) {
		fmt.Println("您所選取的檔案不存在", err_o)
		return
	}

	_, err_n := os.Stat(n)

	if os.IsNotExist(err_n) {
		os.Rename(o, n)
		fmt.Println("檔案重新命名成功")
		return
	}

	fmt.Printf("檔案名稱重複 %v，請改用其它名稱\n", n)
}
