package fileutil

import (
	"io"
	"fmt"
	"os"
	"logger"
	"bufio"
)

const TAG = "FILEUTIL"

func OpenPbmFile(name string) io.Reader {
	filename := fmt.Sprintf("%s.in", name)
	logger.V(TAG, "Open pbm %s", filename)

	file, err := os.Open(filename)
	check(err)
	return file
}

func CreateSubFile(name string, score int) *bufio.Writer {
	filename := fmt.Sprintf("%s_%d.out", name, score)
	logger.V(TAG, "Open sub %s", filename)

	file, err := os.Create(filename)
	check(err)
	return bufio.NewWriter(file)
}

func OpenSubFile(name string) io.Reader {
	filename := fmt.Sprintf("%s.out", name)
	logger.V(TAG, "Open sub %s", filename)

	file, err := os.Open(filename)
	check(err)
	return file
}


func check(e error) {
	if e != nil {
		logger.Wtf(TAG, "oups!!!")
		logger.Wtf(TAG, "error %v", e)
		panic(e)
	}
}
