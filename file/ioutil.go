package file

import (
	"fmt"
	"io/ioutil" // The latest version is not going to support this anymore, please use io instead of ioutil functions
	"os"
	"strings"
)

func IoutilRead(src string) {
	file, err := ioutil.ReadFile(src)
	fmt.Println(err) // <nil>
	HandleErr(err)
	// file is a byte type
	fmt.Printf("%T\n", file)  // []uint8
	fmt.Println(file)         // [65 66 67 68 69 70]
	fmt.Println(string(file)) // ABCDEF
}

func IoutilWrite(src, str string) {
	// The file will be auto-generator if not found and write the new content
	// The content would be replaced if there are any contents
	file := ioutil.WriteFile(src, []byte(str), os.ModePerm)
	fmt.Println(file)
}

func IoutilReadAll(str string) {
	r := strings.NewReader(str) // convert string into io.Reader
	s, err := ioutil.ReadAll(r) // convert to splice []byte
	fmt.Println(err)            // <nil>
	fmt.Println(s)              // []byte
	fmt.Println(string(s))      // string
}

// Only read one level
func IoutilReadDir(src string) {
	f, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Println(err) // <nil>
		return
	}
	// fmt.Printf("%v is %T\n", f, f) // [0xc000188410 0xc0001884e0 0xc0001885b0] is []os.FileInfo
	fmt.Println("Total: ", len(f)) // 3
	for i := 0; i < len(f); i++ {
		fmt.Printf("%d: %s, isFolder: %t\n", i, f[i].Name(), f[i].IsDir())
	}

	/*  os.Stat(), ioutil.ReadDir() return os.FileInfo type
	Name() string return file name or folder name
	Size() int64 return file or folder size
	Mode() FileMode reeturn perms
	ModTime() time.Time return modify time
	IsDir() bool return check the current target is folder
	Sys() interface{}
	*/
}
