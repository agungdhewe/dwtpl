package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/agungdhewe/dwtpl"
)

type PageData struct {
	Nama   string
	Alamat string
}

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

	fmt.Println(page.Config.Title)

	tpl, inmap := page.Data[dwtpl.DeviceMobile]
	if !inmap {
		fmt.Println("tidak ada halaman mobile")
		os.Exit(1)
	}

	data := &PageData{
		Nama:   "Agung",
		Alamat: "Jakarta",
	}

	buff := new(bytes.Buffer)
	err = tpl.Execute(buff, data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(buff.String())

}
