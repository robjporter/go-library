package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/robjporter/go-library/xsecretbox"
)

func main() {
	var fileRaw = "./testdata/index.txt"
	var fileEnc = "./testdata/encrypted.bin"

	msg, err := ioutil.ReadFile(fileRaw)
	if err != nil {
		t.Fatalf("ioutil.ReadFile(%s) failed: %v", fileRaw, err)
	}

	// encrypt plain text file
	err = Seal(fileRaw, fileEnc)
	if err != nil {
		t.Fatalf("Seal() failed: %v", err)
	}

	tmp, err := ioutil.ReadFile(fileEnc)
	if err != nil {
		t.Fatalf("ioutil.ReadFile(%s) failed: %v", fileEnc, err)
	}

	// decrypt encrypted file and compare results
	err = Open(fileEnc, fileRaw)
	if err != nil {
		t.Fatalf("Seal() failed: %v", err)
	}

	tmp, err = ioutil.ReadFile(fileRaw)
	if err != nil {
		t.Fatalf("ioutil.ReadFile(%s) failed: %v", fileRaw, err)
	}

	if !bytes.Equal(tmp, msg) {
		t.Fatal("plainFileTemp != plainFileTest")
	}
}
