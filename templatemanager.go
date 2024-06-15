package dwtpl

import (
	"fmt"
	"html/template"
	"io"
	"log"

	"github.com/agungdhewe/dwpath"
)

type TemplateManager struct {
	logger           *log.Logger
	configuration    Configuration
	pagesDirLocation string
	cachedata        map[string]map[DeviceType]*template.Template
	options          []string
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
func New(config *Configuration, opt ...string) (*TemplateManager, error) {
	var exists bool

	// siapkan template manager
	mgr = &TemplateManager{
		logger:        log.New(log.Writer(), "", 0),
		configuration: *config,
		cachedata:     make(map[string]map[DeviceType]*template.Template),
		options:       opt,
	}

	// cek apakah direktori template ada
	exists, _ = dwpath.IsDirectoryExists(mgr.configuration.Dir)
	if !exists {
		return nil, fmt.Errorf("direktori template %s tidak ditemukan", mgr.configuration.Dir)
	}

	mgr.logger.SetOutput(io.Discard)

	return mgr, nil
}

// GetOptions returns the options stored in the TemplateManager.
//
// No parameters.
// Returns a slice of strings representing the options.
func (mgr *TemplateManager) GetOptions() []string {
	return mgr.options
}

// SetOptions sets the options for the TemplateManager.
//
// It takes a variadic parameter `opt` of type string, which represents the options to be set.
// The function updates the `options` field of the TemplateManager with the provided options.
func (mgr *TemplateManager) SetOptions(opt ...string) {
	mgr.options = opt
}

// SetLogOutput sets the output writer for the logger of the TemplateManager.
//
// It takes a parameter `w` of type `io.Writer`, which represents the writer to
// which the logger output will be redirected.
// The function updates the logger's output writer with the provided writer.
func (mgr *TemplateManager) SetLogOutput(w io.Writer) {
	mgr.logger.SetOutput(w)
}

// GetConfiguration returns the configuration of the TemplateManager.
//
// This function returns a pointer to the Configuration struct stored in the TemplateManager.
// It allows the caller to access and modify the configuration.
//
// Returns:
// - *Configuration: A pointer to the Configuration struct.
func (mgr *TemplateManager) GetConfiguration() *Configuration {
	return &mgr.configuration
}

// GetPage retrieves a page template for the specified page name and device type.
//
// Parameters:
// - pagename: the name of the page to retrieve the template for.
// - device: the device type for which the page template is requested.
//
// Returns:
// - *template.Template: the retrieved page template, or nil if not found.
// - bool: true if the page template exists, false otherwise.
// - error: an error if there was a problem retrieving the page template.
func (mgr *TemplateManager) GetPage(pagename string, device DeviceType) (*template.Template, bool, error) {
	var tpl *template.Template
	var exists bool
	var err error

	exists = false
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
		pagedata, ispage, err = mgr.ParsePageTemplate(pagename, mgr.pagesDirLocation)
		if err != nil {
			report_error(err.Error())
			return nil, false, fmt.Errorf("tidak dapat parse halaman %s", pagename)
		}

		if !ispage {
			return nil, false, fmt.Errorf("struktur pada halaman %s tidak ditemukan", pagename)
		}

		tpl, exists = pagedata[device]
		if !exists {
			report_log("halaman %s untuk device %s tidak ditemukan", pagename, device)
		}

		// apabila configured dengan cache, simpan kembali data ke cache
		if mgr.configuration.Cached {
			mgr.cachedata[pagename] = pagedata
		}

	}
	report_log("ok, sajikan halaman %s untuk device %s", pagename, device)
	return tpl, true, nil

}
