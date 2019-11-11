package common

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"math/big"
	"os"
	"strings"
)

func GetBigInt(str string, base int) *big.Int {
	res, _ := new(big.Int).SetString(str, base)
	return res
}
func ReadZkInfoFromFile(filename string) ([]string, error) {
	vals := readFileWithReadLine(filename)
	if len(vals) != 7 {
		return nil, errors.New("invalid Zk file. Please check again")
	}
	return vals, nil
}
func readFileWithReadLine(path string) (values []string) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	for {
		var buffer bytes.Buffer
		var l []byte
		var isPrefix bool

		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)

			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				break
			}

			// If we're just at the EOF, break
			if err != nil {
				break
			}
		}

		if err == io.EOF {
			break
		}

		line := buffer.String()
		if len(line) == 0 || line == "" || line == " " {
			continue
		}
		line = strings.ReplaceAll(line, "\"", "")
		line = strings.ReplaceAll(line, " ", "")
		//values = append(values, strings.TrimSpace(line))
		strings.TrimSpace(line)

		strs := strings.Split(line, ",")
		values = append(values, strs...)
	}
	if err != io.EOF && err != nil {
		panic(err)
	}
	return values
}
