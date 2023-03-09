package file

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ListFiles(dir string, level int) {
	// level用來記錄當前遞歸的層次，生成帶有層次的空格
	s := "|--"

	for i := 0; i < level; i++ {
		s = "| " + s
	}
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fileInfos {

		filename := dir + "/" + f.Name()
		fmt.Println(s + filename)

		if f.IsDir() {
			// 遞歸調用方法
			ListFiles(filename, level+1)
		}
	}
}
