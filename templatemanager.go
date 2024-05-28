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

// siapkan template manager
// sebelum memanggil ini, harus panggil dwtpl.New() terlebih dahulu
func New(config *Configuration) (*TemplateManager, error) {
	var exists bool

	// siapkan template manager
	mgr = &TemplateManager{
		logger:        log.New(log.Writer(), "", 0),
		configuration: *config,
		cachedata:     make(map[string]map[DeviceType]*template.Template),
	}

	// cek apakah direktori template ada
	exists, _ = dwpath.IsDirectoryExists(mgr.configuration.Dir)
	if !exists {
		return nil, fmt.Errorf("direktori template %s tidak ditemukan", mgr.configuration.Dir)
	}

	mgr.logger.SetOutput(io.Discard)

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

func (mgr *TemplateManager) GetPage(pagename string, device DeviceType) (tpl *template.Template, exists bool, err error) {

	if mgr.configuration.Cached {
		// ambil dari cache
		report_log("cek data %s dari cache", pagename)
		tpl, exists = mgr.GetCachedPage(pagename, device)
		if !exists {
			report_log("data halaman %s tidak ditemukan di cache", pagename)
		}
	}

	if !exists {
		// di cache belum ada, coba cari langsung dari disk
		var pagedata map[DeviceType]*template.Template
		var ispage bool
		report_log("ambil data %s dari disk", pagename)

		// cek direktori halaman
		pagedir := filepath.Join(mgr.pagesDirLocation, pagename)
		exists, err = dwpath.IsDirectoryExists(pagedir)
		if !exists {
			if err != nil {
				report_error("ada kesalahan saat cek direktori %s", pagedir)
				return nil, false, err
			} else {
				report_error("direktori %s tidak ditemukan", pagedir)
				return nil, false, nil
			}
		}

		// mulai parse halaman
		pagedata, ispage, err = mgr.ParsePageTemplate(pagedir)
		if err != nil {
			report_error("tidak dapat parse halaman %s", pagename)
			return nil, false, err // pagedata, exists, error
		}

		if !ispage {
			report_error("struktur pada %s tidak sesuai dengan struktur halaman", pagename)
			return nil, false, fmt.Errorf("struktur pada halaman %s tidak ditemukan", pagename)
		}

		tpl, exists = pagedata[device]
		if !exists {
			report_error("halaman %s untuk device %s tidak ditemukan", pagename, device)
		}

		// apabila configured dengan cache, simpan kembali data ke cache
		if mgr.configuration.Cached {
			mgr.cachedata[pagename] = pagedata
		}

	}
	report_log("ok, sajikan halaman %s untuk device %s", pagename, device)
	return tpl, true, nil

}
