package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/agungdhewe/dwtpl"
)

func main() {

	// rootDir, err := os.Getwd()
	_, filename, _, _ := runtime.Caller(0)
	curdir := filepath.Dir(filename)
	tpldir := filepath.Join(curdir, "template")

	config := &dwtpl.Configuration{
		Dir:    tpldir,
		Cached: true,
	}

	// inisiasi modul
	mgr, err := dwtpl.New(config)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("tidak inisiasi template")
		os.Exit(1)
	}

	// untuk keperluan debug
	mgr.SetLogOutput(log.Writer())

	pagedir := filepath.Join(curdir, "pages")
	err = mgr.CachePages(pagedir)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("tidak bisa cache halaman")
		os.Exit(1)
	}

	page, err := mgr.GetPage("home")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
