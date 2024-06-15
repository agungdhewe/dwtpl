package dwtpl

import (
	"fmt"
	"html/template"
	"path/filepath"
)

// CachePages caches the pages in the specified directory by parsing the page templates and storing them in the cache.
//
// Parameters:
// - dir: The directory where the pages are located.
//
// Returns:
// - err: An error if there was a problem caching the pages.
func (mgr *TemplateManager) CachePages(dir string) (err error) {
	var pages []string

	mgr.pagesDirLocation = dir

	// baca direktori pagedir
	pattern := filepath.Join(mgr.pagesDirLocation, "*")
	pages, err = filepath.Glob(pattern)
	if err != nil {
		report_error(err.Error())
		return fmt.Errorf("tidak dapat mengambil data pattern %s", pattern)
	}

	for _, pagedir := range pages {
		pagename := filepath.Base(pagedir)
		tpldata, ispage, err := mgr.ParsePageTemplate(pagedir)
		if err != nil {
			report_error(err.Error())
			return fmt.Errorf("tidak dapat parse halaman %s", pagename)
		}

		if ispage {
			mgr.cachedata[pagename] = tpldata
		}
	}

	return nil
}

// GetCachedPage retrieves a cached page template for the specified page name and device type.
//
// Parameters:
// - pagename: The name of the page to retrieve the template for.
// - device: The device type for which to retrieve the template.
//
// Returns:
// - tpl: The cached page template, or nil if not found.
// - err: A boolean indicating if an error occurred while retrieving the template.
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
