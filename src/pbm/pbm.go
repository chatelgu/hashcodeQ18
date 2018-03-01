package pbm

import (
	"fileutil"
	"fmt"
	"logger"
)

const TAG = "PBM"


type Pbm struct {
	Name string
}

// constructors

func Parse(name string) Pbm {
	var p Pbm
	p.Name = name

	// open the file
	reader := fileutil.OpenPbmFile(name)

	// read data and construct pbm
	var i, j , k int
	fmt.Fscanf(reader,"%d %d %d", &i, &j , &k)
	logger.D(TAG, "read %d %d %d", i , j, k)
	fmt.Fscanf(reader, "%d", &i)
	logger.D(TAG, "read %d %d %d", i , j, k)
	fmt.Fscanf(reader, "%d %d", &i, &j)
	logger.D(TAG, "read %d %d %d", i , j, k)
	fmt.Fscanf(reader, "%d %d %d", &i, &j, &k)
	logger.D(TAG, "read %d %d %d", i , j, k)

	return p
}

// pretty print

func (p *Pbm) String() string {
	return fmt.Sprintf("Pbm: %s", p.Name)
}