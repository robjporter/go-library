package main

import (
	"github.com/robjporter/go-library/xfastlog"
)

func main() {
	var flog xfastlog.FastLog = xfastlog.NewConsoleLogger("fastlog")
	flog.Info("hello fastlog!!!")
	flog.Warning("hello fastlog!!!")
	flog.Notice("hello fastlog!!!")
	flog.Debug("hello fastlog!!!")
	flog.Error("hello fastlog!!!")
	flog.Fatal("hello fastlog!!!")

	var flog2 xfastlog.FastLog = xfastlog.NewRotateLogger("./fastlog", "fastlog", 1024*1024, 5)
	flog2.Info("hello fastlog!!!")
	flog2.Warningf("hello fastlog!!!")
	flog2.Noticef("hello fastlog!!!")
}
