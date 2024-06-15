package dwtpl

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/agungdhewe/dwpath"
)

// ParsePageTemplate parses the page template for a given page name and data directory.
//
// Parameters:
// - pagename: the name of the page (string)
// - pagesdatadir: the directory where the page data is located (string)
//
// Returns:
// - tpldata: a map of device types to template.Template objects (map[DeviceType]*template.Template)
// - ispage: a boolean indicating whether the page directory exists (bool)
// - err: an error if there was an issue parsing the page template (error)
func (mgr *TemplateManager) ParseTemplate(pagename string, pagesdatadir string) (*PageTemplate, error) {
	var exists bool
	var err error

	// ambil data layout
	layoutfiles, _, err := mgr.GetLayoutFiles(mgr.configuration.Dir)
	if err != nil {
		report_error(err.Error())
		return nil, fmt.Errorf("error saat memuat layout template dari %s", mgr.configuration.Dir)
	}

	// cek direktori halaman
	pagedir := filepath.Join(pagesdatadir, pagename)
	exists, err = dwpath.IsDirectoryExists(pagedir)
	if !exists {
		if err != nil {
			report_error(err.Error())
			return nil, fmt.Errorf("direktori %s tidak ditemukan", pagedir)
		} else {
			report_error("direktori %s tidak ditemukan", pagedir)
			return nil, nil
		}
	}

	// ambil data halaman
	var pageconfig *PageConfig

	pageconfig, err = mgr.GetPageConfig(pagedir)
	if err != nil {
		report_error(err.Error())
		return nil, fmt.Errorf("tidak dapat memuat konfigurasi halaman %s", pagedir)
	}

	tpldata := map[DeviceType]*template.Template{}
	for _, device := range []DeviceType{DeviceMobile, DeviceTablet, DeviceDesktop} {
		var tpl *template.Template
		files := append(pageconfig.layoutfiles[device], layoutfiles[device]...)

		t := template.New(fmt.Sprintf("%s.html", pagename))
		if mgr.options != nil {
			t.Option(mgr.options...)
		}

		tpl, err = t.ParseFiles(files...)
		if err != nil {
			report_error(err.Error())
			return nil, fmt.Errorf("tidak dapat parse file template untuk halaman %s", pagename)
		}
		tpldata[device] = tpl
	}

	pagetemplate := &PageTemplate{
		Config: pageconfig,
		Data:   tpldata,
	}

	return pagetemplate, nil
}
