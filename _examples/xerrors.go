package main

import (
	"fmt"
	"io/ioutil"

	"github.com/robjporter/go-library/xerrors"
)

func main() {
	a := readFile("")
	fmt.Println(a)
	fmt.Println("====================")
	err := xerrors.Cause(a)
	fmt.Println(err)
}

func readFile(r string) error {
	_, err := ioutil.ReadFile(r)
	if err != nil {
		return xerrors.Wrap(err, "read failed")
	}
	return nil
}