package dwtpl

import (
	"path/filepath"
	"text/template"

	"github.com/agungdhewe/dwpath"
)

func (mgr *TemplateManager) ParsePageTemplate(pagedir string) (tpldata map[DeviceType]*template.Template, ispage bool, err error) {
	// ambil data layout
	layoutfiles, _, err := mgr.GetLayoutFiles(mgr.configuration.Dir)
	if err != nil {
		report_error("error saat memuat layout template")
		return nil, false, err
	}

	// ambil data halaman
	var pagefiles map[DeviceType][]string
	pagefiles, ispage, err = mgr.GetLayoutFiles(pagedir)
	if err != nil {
		report_error("error saat memuat layout halaman")
		return nil, false, err //tpldata, ispage, err
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
			report_error("tidak dapat parse file template untuk halaman dari %s", pagedir)
			return nil, false, err
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
		report_error("error saat memuat layout template")
		return nil, false, err
	}

	// cek direktori halaman
	pagedir := filepath.Join(pagesdatadir, pagename)
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

	// ambil data halaman
	var pagefiles map[DeviceType][]string
	pagefiles, ispage, err = mgr.GetLayoutFiles(pagedir)
	if err != nil {
		report_error("error saat memuat layout halaman")
		return nil, false, err //tpldata, ispage, err
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
			report_error("tidak dapat parse file template untuk halaman %s", pagename)
			return nil, false, err
		}
		tpldata[device] = tpl
	}

	return tpldata, true, nil
}
