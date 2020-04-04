package main

import (
	"fmt"
	"github.com/robjporter/go-library/xenv"
)

func main() {
	env := xenv.New()
	fmt.Println("SIZE: ", env.Size())
	fmt.Println(env.GetString("GOPATH"))
	fmt.Println("SIZE: ", env.Size())
	env.AddString("TEST", "TESTING")
	fmt.Println("SIZE: ", env.Size())
	fmt.Println("DATA: ", env.GetAll())
	fmt.Println("")
	fmt.Println("SYSTEM ENVIRONMENTS *******************************************************")
	fmt.Println("ENVIRONMENT GO PATH BIN:          >", env.GOPATHBIN())
	fmt.Println("ENVIRONMENT GO PATH:              >", env.GOPATH())
	fmt.Println("ENVIRONMENT PATH SEPARATOR:       >", env.PathSeparator())
	fmt.Println("ENVIRONMENT LIST SEPARATOR:       >", env.ListSeparator())
	fmt.Println("ENVIRONMENT IS COMPILED:          >", env.IsCompiled())
	fmt.Println("ENVIRONMENT BUILD DEBUG:          >", env.BuildDebug())
	fmt.Println("ENVIRONMENT CHECK ARCH:           >", env.CheckArchitecture())
	fmt.Println("ENVIRONMENT BUILD STAMP:          >", env.BuildStamp())
	fmt.Println("ENVIRONMENT COMPILER:             >", env.Compiler())
	fmt.Println("ENVIRONMENT GO ARCH:              >", env.GOARCH())
	fmt.Println("ENVIRONMENT GO OS:                >", env.GOOS())
	fmt.Println("ENVIRONMENT GO ROOT:              >", env.GOROOT())
	fmt.Println("ENVIRONMENT GO VERSION:           >", env.GOVER())
	fmt.Println("ENVIRONMENT NUMBER CPU:           >", env.NumCPU())
	fmt.Println("ENVIRONMENT FORMATTED TIME:       >", env.GetFormattedTime())
	fmt.Println("ENVIRONMENT USERNAME:             >", env.GetUsername())
}
