package dwtpl

import (
	"path/filepath"
	"text/template"
)

func (mgr *TemplateManager) CachePages(dir string) (err error) {
	var pages []string

	mgr.pagesDirLocation = dir

	// baca direktori pagedir
	pattern := filepath.Join(mgr.pagesDirLocation, "*")
	pages, err = filepath.Glob(pattern)
	if err != nil {
		report_error("tidak dapat mengambil data pattern %s", pattern)
		return err
	}

	for _, pagedir := range pages {
		pagename := filepath.Base(pagedir)
		tpldata, ispage, err := mgr.ParsePageTemplate(pagedir)
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

func (mgr *TemplateManager) GetCachedPage(pagename string, device DeviceType) (tpl *template.Template, err bool) {
	pagedata, pageinmap := mgr.cachedata[pagename]
	if !pageinmap {
		return nil, false
	}

	var deviceinmap bool
	tpl, deviceinmap = pagedata[device]
	if !deviceinmap {
		return nil, false
	}

	return tpl, true
}
