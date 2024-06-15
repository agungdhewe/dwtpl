package dwtpl

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"path/filepath"

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

// New initializes a new TemplateManager with the given Configuration.
//
// Parameters:
// - config: a pointer to a Configuration struct.
//
// Returns:
// - a pointer to a TemplateManager struct, or nil if the template directory is not found.
// - an error if the template directory is not found.
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

// SetLogOutput sets the output destination for the logger.
//
// Parameters:
// - w: an io.Writer interface to set as the output destination.
func (mgr *TemplateManager) SetLogOutput(w io.Writer) {
	mgr.logger.SetOutput(w)
}

// GetConfiguration returns the Configuration object of the TemplateManager.
//
// It does not take any parameters.
// It returns a pointer to a Configuration object.
func (mgr *TemplateManager) GetConfiguration() *Configuration {
	return &mgr.configuration
}

// GetPage retrieves a template for a given page and device type.
//
// Parameters:
// - pagename: the name of the page to retrieve.
// - device: the device type for which to retrieve the template.
//
// Returns:
// - tpl: the template for the given page and device type.
// - exists: a boolean indicating whether the template exists.
// - err: an error if any occurred during the retrieval process.
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
				report_error(err.Error())
				return nil, false, fmt.Errorf("tidak dapat cek direktori %s", pagedir)
			} else {
				report_error("direktori %s tidak ditemukan", pagedir)
				return nil, false, nil
			}
		}

		// mulai parse halaman
		pagedata, ispage, err = mgr.ParsePageTemplate(pagedir)
		if err != nil {
			report_error(err.Error())
			return nil, false, fmt.Errorf("tidak dapat parse halaman %s", pagename)
		}

		if !ispage {
			return nil, false, fmt.Errorf("struktur pada halaman %s tidak ditemukan", pagename)
		}

		tpl, exists = pagedata[device]
		if !exists {
			return nil, false, fmt.Errorf("halaman %s untuk device %s tidak ditemukan", pagename, device)
		}

		// apabila configured dengan cache, simpan kembali data ke cache
		if mgr.configuration.Cached {
			mgr.cachedata[pagename] = pagedata
		}

	}

	// semua ok
	report_log("ok, sajikan halaman %s untuk device %s", pagename, device)
	return tpl, true, nil

}
