package file

import (
	"bufio"
	"fmt"
	"os"
)

func Bufio(src, addStr string) {
	file, err := os.OpenFile(src, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	HandleErr(err)
	defer file.Close()

	p := make([]byte, 1024)
	file.Read(p)

	f := bufio.NewWriter(file)
	n, err := f.WriteString("\n" + addStr)
	fmt.Println(err)
	fmt.Println(n)

	f.Flush() // 刷新緩衝區，如果沒有這一行，被寫入的文字只會存在緩衝區
}
