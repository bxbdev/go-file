package file

import (
	"fmt"
	// "io"
	"os"
)

func WriteFile(f string, b *[]byte, s string) {

	// Open() We cannot use this method here, because it would not be allowed to open the file
	// Error message below:
	// err write ../go-file/assets/docs/test.txt: bad file descriptor

	// We need to use OpenFile() this method to make the parameters with permissions
	// parameters in OpenFile()
	// os.O_TRUNC means if the file is existed will clean all contents and replace new content
	// os.O_WRONLY means allow to write contents into file
	// os.O_CREATE means if not existed then create new file
	// os.O_APPEND means to add contents from the last word

	file, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	HandleErr(err)
	defer file.Close()

	// Write contents into file
	// bs := []byte{65, 66, 67, 68, 69, 70} // A,B,C,D,E,F
	// bs := []byte{97, 98, 99, 100} // abcd
	bs := *b
	n, err := file.Write(bs)
	HandleErr(err)
	fmt.Println(n) // 6

	// make new line
	file.WriteString("\n")
	// write contents by strings
	v, err := file.WriteString(s)
	HandleErr(err)
	fmt.Println(v)

	file.WriteString("\n")

	// convert string to byte
	stb, err := file.Write([]byte("New content"))
	fmt.Println(stb)
	HandleErr(err)
}
