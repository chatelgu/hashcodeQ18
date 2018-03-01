package sub

import (
	"fileutil"
	"fmt"
)

const TAG = "SUB"


type Sub struct {

}

func (s *Sub) ToFile(name string) {
	score := 0
	writer := fileutil.CreateSubFile(name, score)

	fmt.Fprint(writer, "0")
	writer.Flush()
}