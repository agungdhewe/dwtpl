package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/agungdhewe/dwtpl"
)

func main() {
	var err error
	var mgr *dwtpl.TemplateManager
	// var files []string

	// rootDir, err := os.Getwd()
	_, filename, _, _ := runtime.Caller(0)
	curdir := filepath.Dir(filename)
	tpldir := filepath.Join(curdir, "template")

	// inisiasi modul
	dwtpl.New(&dwtpl.TemplateConfig{
		Dir:    tpldir,
		Cached: true,
	})

	// buat template manager
	mgr, err = dwtpl.NewTemplateManager()
	if err != nil {
		fmt.Println(err.Error())
	}

	// coba ambil file layout
	// files, err = mgr.GetLayoutFiles(dwtpl.GetConfig().Dir, dwtpl.DeviceMobile)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// // fmt.Println("layout html template list")
	// // fmt.Println(files)

	var files *map[dwtpl.DeviceType][]string
	var pagedir = filepath.Join(curdir, "pages", "home")
	files, err = mgr.GetLayoutFiles(pagedir)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(&files)
	// coba parsePage

	/*
		pagename := "home"
		pagedir := filepath.Join(curdir, "pages")
		var tpl *template.Template
		tpl, err = mgr.ParseTemplate(pagename, pagedir, dwtpl.DeviceMobile)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(tpl)
	*/
}
