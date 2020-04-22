package main

import (
	"fmt"
	"io/ioutil"

	"github.com/robjporter/go-library/xerrors"
)

func main() {
	a := readFile("")
	fmt.Println(a)
	err := xerrors.Cause(a)
	fmt.Println(err)
}

func readFile(r string) error {
	_, err := ioutil.ReadAll(r)
	if err != nil {
		return xerrors.Wrap(err, "read failed")
	}
}