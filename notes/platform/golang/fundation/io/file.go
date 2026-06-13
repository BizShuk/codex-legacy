package io

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// [Hint]: file permission flags
//    - O_RDONLY, Read-Only Mode.
//    - O_WRONLY. Opens file in Write Only mode.
//    - O_RDWR, Open the file in Read-Write mode.
//    - O_APPEND, Appends data to the file when writing.
//    - O_TRUNC, Truncate the file when opened, where applicable.
//    - O_SYNC, Open the file for synchronous I/O.
//    - O_CREATE, Create a file if it does not exist.

func OpenFile(filepath string) {
	file, err := os.Open(filepath /* , permission flags, file mode*/)

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("File close failure with err:", err)
		}
	}()

	if err != nil {
		logrus.Fatalln("open file failed with error:", err)
	}
	BufferReader_ReadString(file)
}

func FileRead(file *os.File) {
	data := make([]byte, 50)
	count, err := file.Read(data)
	if err != nil {
		logrus.Fatalln("Read file data failed with error:", err)
	}

	fmt.Printf("Read %d bytes: %q", count, data[:count])
}

func FileRead_simlpe(s string) {
	dat, err := os.ReadFile(s)
	if err != nil {
		logrus.Fatalln("Read file data failed with error:", err)
	}
	fmt.Print(string(dat))
}

func BufferReader_peak(file *os.File) {
	reader := bufio.NewReader(file)
	b4, err := reader.Peek(5)
	if err != nil {
		logrus.Fatalln("Read file data failed with error:", err)
	}
	fmt.Printf("%v bytes: %s\n", len(b4), string(b4))
}

func BufferReader_ReadLine(file *os.File) {
	reader := bufio.NewReader(file)
	// [Hint]: If one line is too long, isPrefix will be set util last fragement of the line.
	bytes, isPrefix, err := reader.ReadLine()
	if err != nil {
		logrus.Fatalln("BufferRead ReadLine failed with error:", err)
	}

	fmt.Printf("BufferRead ReadLine: %v bytes, isPrefix %v : %s\n", len(bytes), isPrefix, string(bytes))
}

func BufferReader_ReadString(file *os.File) {
	delimiter := '\n' // [issue]: how to assign a char to type byte
	reader := bufio.NewReader(file)
	s, err := reader.ReadString('\n')
	if err != nil {
		logrus.Fatalln("Read file data failed with error:", err)
	}
	fmt.Printf("BufferReader ReadString: delimiter %v :  %s\n", delimiter, s)
}
