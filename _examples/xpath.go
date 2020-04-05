package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/robjporter/go-library/xpath"
)

func main() {
	p := xpath.New()

	fmt.Println("")
	fmt.Println("XPATH *******************************************************")
	dir, err := filepath.Abs("./")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	p.UpdatePath(dir)
	fmt.Println("Path:                          >", dir)
	fmt.Println("Splitpath:                     >", p.SplitPath())
	fmt.Println("Get Path:                      >", p.GetPath())
	fmt.Println("Get File MD5:                  >", p.GetFileMd5("./main.go"))
	fmt.Println("Make Directory:                >", p.MkDirSpecificMode("test", os.FileMode.IsDir(true)))
	fmt.Println("Make Directory 2:              >", p.MkDir("test2"))
	fmt.Println("Is Dir Directory:              >", p.IsDir("test"))
	fmt.Println("Is Dir Directory 2:            >", p.IsDir("test2"))
	fmt.Println("Is Exist Directory:            >", p.IsExist("test"))
	fmt.Println("Is Exist Directory 2:          >", p.IsExist("test2"))
	fmt.Println("Parentpath:                    >", p.ParentPath())
	fmt.Println("Relativepath:                  >", p.RelativePath(p.ParentPath()))
	fmt.Println("BaseName:                      >", p.BaseName())
	fmt.Println("ListFiles:                     >", p.ListFilesRecursive("", "", false))
	tmp := p.ListFilesRecursive("", "", false)
	for i := 0; i < len(tmp); i++ {
		pather := dir + "/" + tmp[i]
		tmp2, _ := p.FileMode(pather)
		tmp3, _ := p.FileSize(pather)
		fmt.Println("ListFilesInfo:                 >", pather, " - ", tmp2)
		fmt.Println("ListFilesSize:                 >", pather, " - ", tmp3)
	}
}