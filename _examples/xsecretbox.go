package main

import (
	"log"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/robjporter/go-library/xsecretbox"
)

func main() {
	tmpdir, err := ioutil.TempDir("", "testdata")
	if err != nil {
		log.Fatalf("ioutil.TempDir() failed: %v", err)
	}
	defer os.RemoveAll(tmpdir)

	plainFileTest := filepath.Join("./", "test.txt")
	plainFileTemp := filepath.Join("testdata", "test.txt")
	cryptFileTemp := filepath.Join("testdata", "test.bin")

	//testPass := "test"

	msg, err := ioutil.ReadFile(plainFileTest)
	if err != nil {
		log.Fatalf("ioutil.ReadFile(%s) failed: %v", plainFileTest, err)
	}

	// encrypt plain text file
	err = xsecretbox.Seal(plainFileTest, cryptFileTemp)
	if err != nil {
		log.Fatalf("Seal() failed: %v", err)
	}
	tmp, err := ioutil.ReadFile(cryptFileTemp)
	if err != nil {
		log.Fatalf("ioutil.ReadFile(%s) failed: %v", cryptFileTemp, err)
	}

	// encrypting to existing file should file
	err = xsecretbox.Seal(plainFileTest, cryptFileTemp)
	if err == nil {
		log.Fatal("Seal() should fail")
	}

	// decrypt encrypted file and compare results
	err = xsecretbox.Open(cryptFileTemp, plainFileTemp)
	if err != nil {
		log.Fatalf("Seal() failed: %v", err)
	}
	tmp, err = ioutil.ReadFile(plainFileTemp)
	if err != nil {
		log.Fatalf("ioutil.ReadFile(%s) failed: %v", plainFileTemp, err)
	}
	if !bytes.Equal(tmp, msg) {
		log.Fatal("plainFileTemp != plainFileTest")
	}

	// decrypting to existing file should fail
	err = xsecretbox.Open(cryptFileTemp, plainFileTemp)
	if err == nil {
		log.Fatal("Open() shoudl fail")
	}

	if err := os.Remove(plainFileTemp); err != nil {
		log.Fatalf("os.Remove(%s) failed: %v", plainFileTemp, err)
	}

	// decrypting with wrong passphrase should fail
	//testPass = "fail"
	err = xsecretbox.Open(cryptFileTemp, plainFileTemp)
	if err == nil {
		log.Fatalf("Open() with wrong passphrase should fail")
	}
}
