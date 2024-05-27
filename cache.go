package dwtpl

import (
	"path/filepath"
	"text/template"
)

func (mgr *TemplateManager) CachePages(dir string) error {
	mgr.pagesDirLocation = dir

	// baca direktori pagedir
	pattern := filepath.Join(mgr.pagesDirLocation, "*")
	pages, err := filepath.Glob(pattern)
	if err != nil {
		report_error("tidak dapat mengambil data pattern %s", pattern)
		return err
	}

	for _, pagedir := range pages {
		pagename := filepath.Base(pagedir)
		tpldata, ispage, err := mgr.ParsePageTemplate(pagename, filepath.Join(pagedir, ".."))
		if err != nil {
			report_error("error saat parse halaman %s", pagename)
			return err
		}

		if ispage {
			mgr.cachedata[pagename] = tpldata
		}
	}

	return nil
}

func (mgr *TemplateManager) GetCachedPage(pagename string, device DeviceType) (*template.Template, bool) {
	pagedata, pageinmap := mgr.cachedata[pagename]
	if !pageinmap {
		return nil, false
	}

	tpl, deviceinmap := pagedata[device]
	if !deviceinmap {
		return nil, false
	}

	return tpl, true
}
