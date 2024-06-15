package dwtpl

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/agungdhewe/dwpath"
)

// ParsePageTemplate parses the page template from the specified directory.
//
// Parameter:
//
//	pagedir - the directory path where the page template is located.
//
// Return type(s):
//
//	tpldata - a map of DeviceType to template.Template containing parsed templates for different devices.
//	ispage - a boolean indicating if the directory is a valid page directory.
//	err - an error indicating any issues encountered during parsing.
func (mgr *TemplateManager) ParsePageTemplate(pagedir string) (tpldata map[DeviceType]*template.Template, ispage bool, err error) {
	// ambil data layout
	layoutfiles, _, err := mgr.GetLayoutFiles(mgr.configuration.Dir)
	if err != nil {
		report_error(err.Error())
		return nil, false, fmt.Errorf("tidak dapat memuat layout template dari %s", mgr.configuration.Dir)
	}

	// ambil data halaman
	var pagefiles map[DeviceType][]string
	pagefiles, ispage, err = mgr.GetLayoutFiles(pagedir)
	if err != nil {
		report_error(err.Error())
		return nil, false, fmt.Errorf("tidak dapat memuat layout halaman dari %s", pagedir)
	}

	if !ispage {
		report_log("%s bukan direktori halaman", pagedir)
		return nil, false, nil
	}

	tpldata = map[DeviceType]*template.Template{}
	for _, device := range []DeviceType{DeviceMobile, DeviceTablet, DeviceDesktop} {
		var tpl *template.Template
		files := append(pagefiles[device], layoutfiles[device]...)

		t := template.New(fmt.Sprintf("%s.html", pagename))
		if mgr.options != nil {
			t.Option(mgr.options...)
		}

		tpl, err = t.ParseFiles(files...)
		if err != nil {
			report_error(err.Error())
			return nil, false, fmt.Errorf("tidak dapat parse file template untuk halaman dari %s", pagedir)
		}
		tpldata[device] = tpl
	}

	return tpldata, true, nil
}

func (mgr *TemplateManager) ParsePageTemplateX(pagename string, pagesdatadir string) (tpldata map[DeviceType]*template.Template, ispage bool, err error) {
	var exists bool

	// ambil data layout
	layoutfiles, _, err := mgr.GetLayoutFiles(mgr.configuration.Dir)
	if err != nil {
		report_error(err.Error())
		return nil, false, fmt.Errorf("tidak dapat memuat layout template dari %s", mgr.configuration.Dir)
	}

	// cek direktori halaman
	pagedir := filepath.Join(pagesdatadir, pagename)
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

	// ambil data halaman
	var pagefiles map[DeviceType][]string
	pagefiles, ispage, err = mgr.GetLayoutFiles(pagedir)
	if err != nil {
		report_error(err.Error())
		return nil, false, fmt.Errorf("tidak dapat memuat layout halaman dari %s", pagedir)
	}

	if !ispage {
		report_log("%s bukan direktori halaman", pagedir)
		return nil, false, nil
	}

	tpldata = map[DeviceType]*template.Template{}

	for _, device := range []DeviceType{DeviceMobile, DeviceTablet, DeviceDesktop} {
		var tpl *template.Template
		files := append(pagefiles[device], layoutfiles[device]...)
		tpl, err = template.ParseFiles(files...)
		if err != nil {
			report_error(err.Error())
			return nil, false, fmt.Errorf("tidak dapat parse file template untuk halaman %s", pagename)
		}
		tpldata[device] = tpl
	}

	return tpldata, true, nil
}
