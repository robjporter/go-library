package main

import (
	"fmt"
	"github.com/robjporter/go-library/xenv"
)

func main() {
	env := xenv.New()
	fmt.Println("SIZE: ", env.Size())
	fmt.Println(env.Get("GOPATH"))
	fmt.Println("SIZE: ", env.Size())
	env.AddString("TEST", "TESTING")
	fmt.Println("SIZE: ", env.Size())
	fmt.Println("DATA: ", env.GetAll())
	fmt.Println("")
	fmt.Println("SYSTEM ENVIRONMENTS *******************************************************")
	fmt.Println("ENVIRONMENT GO PATH BIN:          >", xenv.GOPATHBIN())
	fmt.Println("ENVIRONMENT GO PATH:              >", xenv.GOPATH())
	fmt.Println("ENVIRONMENT PATH SEPARATOR:       >", xenv.PathSeparator())
	fmt.Println("ENVIRONMENT LIST SEPARATOR:       >", xenv.ListSeparator())
	fmt.Println("ENVIRONMENT IS COMPILED:          >", xenv.IsCompiled())
	fmt.Println("ENVIRONMENT BUILD DEBUG:          >", xenv.BuildDebug())
	fmt.Println("ENVIRONMENT CHECK ARCH:           >", xenv.CheckArchitecture())
	fmt.Println("ENVIRONMENT BUILD STAMP:          >", xenv.BuildStamp())
	fmt.Println("ENVIRONMENT BUILD HOST:           >", xenv.BuildHost())
	fmt.Println("ENVIRONMENT COMPILER:             >", xenv.Compiler())
	fmt.Println("ENVIRONMENT GO ARCH:              >", xenv.GOARCH())
	fmt.Println("ENVIRONMENT GO OS:                >", xenv.GOOS())
	fmt.Println("ENVIRONMENT GO ROOT:              >", xenv.GOROOT())
	fmt.Println("ENVIRONMENT GO VERSION:           >", xenv.GOVER())
	fmt.Println("ENVIRONMENT NUMBER CPU:           >", xenv.NumCPU())
	fmt.Println("ENVIRONMENT FORMATTED TIME:       >", xenv.GetFormattedTime())
	fmt.Println("ENVIRONMENT USERNAME:             >", xenv.GetUsername())
	fmt.Println("ENVIRONMENT VARIABLES:            >", xenv.GetAllEnvironment())
}
