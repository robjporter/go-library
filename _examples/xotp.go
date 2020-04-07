package main

import (
	"bytes"
	"fmt"

	"github.com/robjporter/go-library/xotp"
)

const (
	DIGITS       = 6
	INIT_COUNTER = 141215332523
	INIT_SECONDS = 0
	TIME_STEP    = 1
)

func main() {
	password := "Pa$$w0rd"
	useTime := false
	numEntries := 5

	buf := bytes.NewBufferString(password)

	var hotp xotp.Generator
	if useTime {
		hotp = xotp.NewTOTPGenerator(buf.Bytes(), INIT_SECONDS, TIME_STEP, DIGITS)
	} else {
		hotp = xotp.NewHOTPGenerator(buf.Bytes(), INIT_COUNTER, DIGITS)
	}

	for i := 0; i < numEntries; i++ {
		fmt.Println("OTP: ", hotp.Generate())
	}
}
