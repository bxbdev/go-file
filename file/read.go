package file

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(f string) {
	// Opened file
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	// Closed connection
	defer file.Close()

	// Read file contents
	// Only read the 4 bytes and 4 alphabets
	// bs := make([]byte, 4, 4)
	// Read the all contents
	bs := make([]byte, 1024)

	/*
		// first time read
		n, err := file.Read(bs)
		fmt.Println(err)        // <nil>
		fmt.Println(n)          // 4
		fmt.Println(bs)         // [97 98 99 100]
		fmt.Println(string(bs)) // abcd

		// second time read
		n, err = file.Read(bs)
		fmt.Println(err)        // <nil>
		fmt.Println(n)          // 4
		fmt.Println(bs)         // [101 102 103 104]
		fmt.Println(string(bs)) // efgh

		// third time read
		n, err = file.Read(bs)
		fmt.Println(err)        // <nil>
		fmt.Println(n)          // 2
		fmt.Println(bs)         // [105 106 103 104]
		fmt.Println(string(bs)) // ijgh

		// fourth time read
		n, err = file.Read(bs)
		fmt.Println(err) // EOF, means no more bytes read
		fmt.Println(n)   // 0
	*/

	fmt.Println("檔案內容如下：")
	n := -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("檔案讀取完畢")
			return
		}
		fmt.Println(string(bs[:n]))
	}

	//	if err != nil {
	//		fmt.Println("err", err)
	//		return
	//	}
}
