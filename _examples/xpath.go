package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/robjporter/go-library/xpath"
)

func main() {
	p,_ := xpath.New("")
	fmt.Println("XPATH *******************************************************")
	dir, err := filepath.Abs("./")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	p.UpdatePath(dir)

	fmt.Println("Path:                          >", p.GetRawPath())
	fmt.Println("Splitpath:                     >", p.SplitPath())
	fmt.Println("Get Path:                      >", p.GetPath())
	fmt.Println("Get File MD5:                  >", p.GetFileMd5("xas.go"))
	fmt.Println("Make Directory:                >", p.MkDirSpecificMode("test", os.ModeDir))
	fmt.Println("Make Directory 2:              >", p.MkDir("test2"))
	t, _ := p.IsDir("test")
	fmt.Println("Is Dir Directory:              >", t)
	t, _ = p.IsDir("test2")
	fmt.Println("Is Dir Directory 2:            >", t)
	t, _ = p.IsExist("test")
	fmt.Println("Is Exist Directory:            >", t)
	t, _ = p.IsExist("test2")
	fmt.Println("Is Exist Directory 2:          >", t)
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
	fmt.Println("XPATH - RUN *************************************************")
	fmt.Println(p.Run("ls",true))
	fmt.Println("XPATH - ListAll *********************************************")
	fmt.Println(p.ListAll())
	fmt.Println("XPATH - ListFiles *******************************************")
	fmt.Println(p.ListFiles())
	fmt.Println("XPATH - ListFilesAll ****************************************")
	fmt.Println(p.ListFilesAll())
	fmt.Println("XPATH - ListFolders *****************************************")
	fmt.Println(p.ListFolders())
	fmt.Println("XPATH - ListFoldersAll **************************************")
	fmt.Println(p.ListFoldersAll())
	fmt.Println("XPATH - istFilesType ****************************************")
	fmt.Println(p.ListFilesType("bak"))
	fmt.Println("XPATH *******************************************************")
}
