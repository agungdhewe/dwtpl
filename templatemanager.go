package dwtpl

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/agungdhewe/dwpath"
)

type TemplateManager struct {
}

type Layout struct {
	Name   string `yaml:"name"`
	Device struct {
		Mobile  []string `yaml:"mobile"`
		Tablet  []string `yaml:"tablet"`
		Desktop []string `yaml:"desktop"`
	} `yaml:"device"`
}

var mgr *TemplateManager

// siapkan template manager
// sebelum memanggil ini, harus panggil dwtpl.New() terlebih dahulu
func NewTemplateManager() (*TemplateManager, error) {
	var exists bool

	// cek apakah modul template sudah diisiasi
	// pertama harus panggil dwtpl.New(conf) di sebelum menggunakan TemplateManager
	if cfg == nil {
		return nil, fmt.Errorf("template belum diinisiasi. sebelum menggunakan template manager inisiasi dengan dwtpl.New(config)")
	}

	// siapkan template manager
	mgr = &TemplateManager{}

	// cek apakah direktori template ada
	exists, _ = dwpath.IsDirectoryExists(cfg.Dir)
	if !exists {
		return nil, fmt.Errorf("direktori template %s tidak ditemukan", cfg.Dir)
	}

	return mgr, nil
}

// mengambil daftar file-file di suatu direktori yang akan digunakan untuk melayout tampilan
// berdasar file konfigurasi xxx.yml pada direktori tersebut
func (mgr *TemplateManager) GetLayoutFiles(dir string) (*map[DeviceType][]string, error) {
	var err error

	// siapkan untuk membaca data layout
	basename := filepath.Base(dir)
	ymllayoutfile := fmt.Sprintf("%s.yml", basename)
	ymllayoutpath := filepath.Join(dir, ymllayoutfile)

	// cek file configurasi yml
	exists, _, _ := dwpath.IsFileExists(ymllayoutpath)
	if !exists {
		return nil, fmt.Errorf("file %s layout tidak ditemukan", ymllayoutpath)
	}

	// baca konfigurasi
	layoutconfig := &Layout{}
	err = readLayoutConfigYml(ymllayoutpath, layoutconfig)
	if err != nil {
		return nil, err
	}

	// ambil daftar file sesuai device yang didefinisikan
	var files = make(map[DeviceType][]string)
	files[DeviceMobile] = layoutconfig.Device.Mobile
	files[DeviceTablet] = layoutconfig.Device.Tablet
	files[DeviceDesktop] = layoutconfig.Device.Desktop

	return &files, nil
}

// mengambil data halaman yang telah di parsing
// sesuai template yang dimaksud template dari suatu direktori
func (mgr *TemplateManager) ParseTemplate(pagename string, dir string, device DeviceType) (*template.Template, error) {
	var err error
	var layoutfiles *map[DeviceType][]string
	var pagefiles *map[DeviceType][]string
	// var files []string
	// var fpath string
	// var tpl *template.Template

	// ambil base layout template
	layoutfiles, err = mgr.GetLayoutFiles(GetConfig().Dir)
	if err != nil {
		return nil, err
	}

	// ambil page layout file
	pagedir := filepath.Join(dir, pagename)
	pagefiles, err = mgr.GetLayoutFiles(pagedir)
	if err != nil {
		return nil, err
	}

	fmt.Println(*layoutfiles)
	fmt.Println(*pagefiles)

	return nil, nil
	// files = []string{}

	// // ambil semua list di filelayout, bentuk jadi path
	// for _, filename := range layoutfiles {
	// 	fpath = filepath.Join(GetConfig().Dir, filename)
	// 	files = append(files, fpath)
	// }

	// // ambil semua list di filelayout, bentuk jadi path
	// for _, filename := range pagefiles {
	// 	fpath = filepath.Join(pagedir, filename)
	// 	files = append(files, fpath)
	// }

	// tpl, err = template.ParseFiles(files...)
	// if err != nil {
	// 	return nil, err
	// }

	// return tpl, nil

}

func (mgr *TemplateManager) CachePages(dir string) error {
	var err error
	var pages []string
	var pagename string

	pattern := filepath.Join(dir, "*")
	pages, err = filepath.Glob(pattern)
	if err != nil {
		return err
	}

	for _, pagedir := range pages {
		pagename = filepath.Base(pagedir)
		mgr.ParseTemplate(pagename, pagedir, DeviceMobile)
		mgr.ParseTemplate(pagename, pagedir, DeviceTablet)
		mgr.ParseTemplate(pagename, pagedir, DeviceDesktop)
	}

	return nil
}
