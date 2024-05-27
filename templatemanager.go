package dwtpl

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"text/template"

	"github.com/agungdhewe/dwpath"
)

type TemplateManager struct {
	logger           *log.Logger
	configuration    Configuration
	pagesDirLocation string
	cachedata        map[string]map[DeviceType]*template.Template
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
func New(config *Configuration) (*TemplateManager, error) {
	var exists bool

	// siapkan template manager
	mgr = &TemplateManager{
		logger:        log.New(log.Writer(), "", log.Lmicroseconds|log.Lshortfile),
		configuration: *config,
		cachedata:     make(map[string]map[DeviceType]*template.Template),
	}

	// cek apakah direktori template ada
	exists, _ = dwpath.IsDirectoryExists(mgr.configuration.Dir)
	if !exists {
		return nil, fmt.Errorf("direktori template %s tidak ditemukan", mgr.configuration.Dir)
	}

	return mgr, nil
}

func (mgr *TemplateManager) Ready() {
}

func (mgr *TemplateManager) SetLogOutput(w io.Writer) {
	mgr.logger.SetOutput(w)
}

func (mgr *TemplateManager) GetConfiguration() *Configuration {
	return &mgr.configuration
}

// mengambil daftar file-file di suatu direktori yang akan digunakan untuk melayout tampilan
// berdasar file konfigurasi xxx.yml pada direktori tersebut
func (mgr *TemplateManager) GetLayoutFiles(dir string) (map[DeviceType][]string, bool, error) {
	var err error
	var exists bool

	// siapkan untuk membaca data layout
	basename := filepath.Base(dir)
	ymllayoutfile := fmt.Sprintf("%s.yml", basename)
	ymllayoutpath := filepath.Join(dir, ymllayoutfile)

	// cek file configurasi yml
	exists, _, err = dwpath.IsFileExists(ymllayoutpath)
	if err != nil {
		return nil, false, err
	}

	// kalau file yml tidak ada, berarti bukan direktori halaman
	if !exists {
		return nil, false, nil
	}

	// baca konfigurasi
	layoutconfig := &Layout{}
	err = readLayoutConfigYml(ymllayoutpath, layoutconfig)
	if err != nil {
		return nil, false, err
	}

	// ambil daftar file sesuai device yang didefinisikan
	var files = make(map[DeviceType][]string)
	files[DeviceMobile] = layoutconfig.Device.Mobile
	files[DeviceTablet] = layoutconfig.Device.Tablet
	files[DeviceDesktop] = layoutconfig.Device.Desktop

	return files, true, nil
}

// mengambil data halaman yang telah di parsing
// sesuai template yang dimaksud template dari suatu direktori
func (mgr *TemplateManager) ParsezzzzTemplate(pagename string, dir string, device DeviceType) (*template.Template, error) {
	var err error
	var layoutfiles *map[DeviceType][]string
	var pagefiles *map[DeviceType][]string
	// var files []string
	// var fpath string
	// var tpl *template.Template

	// ambil base layout template
	layoutfiles, err = mgr.GetLayoutFiles(mgr.configuration.Dir)
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
