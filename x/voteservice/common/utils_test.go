package common

import (
	"fmt"
	"testing"
)

func TestReadZkInfoFromFile(t *testing.T) {
	filename := "./testfiles/voter_2.txt"
	zkInfo, err := ReadZkInfoFromFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(zkInfo)

}
