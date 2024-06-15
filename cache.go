package dwtpl

import (
	"fmt"
	"path/filepath"
)

// CachePages caches the pages in the specified directory by parsing the page templates and storing them in the cache.
//
// Parameters:
// - dir: The directory containing the pages to cache.
//
// Returns:
// - error: An error if there was a problem parsing the page templates or if the pattern could not be obtained.
func (mgr *TemplateManager) CachePages(dir string) error {
	mgr.pagesDirLocation = dir

	// baca direktori pagedir
	pattern := filepath.Join(mgr.pagesDirLocation, "*")
	pages, err := filepath.Glob(pattern)
	if err != nil {
		report_error(err.Error())
		return fmt.Errorf("tidak dapat mengambil data pattern %s", pattern)
	}

	for _, pagedir := range pages {
		pagename := filepath.Base(pagedir)
		pagetemplate, err := mgr.ParseTemplate(pagename, filepath.Join(pagedir, ".."))
		if err != nil {
			report_error(err.Error())
			return fmt.Errorf("tidak dapat parse halaman %s", pagename)
		}

		mgr.cachedpages[pagename] = pagetemplate
	}

	return nil
}

// GetCachedPage retrieves a cached page template based on the page name and device type.
//
// Parameters:
// - pagename: The name of the page to retrieve.
// - device: The type of device for which the page template is needed.
// Returns:
// - *template.Template: The cached page template.
// - bool: A boolean indicating if the page template was found in the cache.
func (mgr *TemplateManager) GetCachedPage(pagename string) (*PageTemplate, bool) {
	pagetemplate, pageinmap := mgr.cachedpages[pagename]
	if !pageinmap {
		return nil, false
	}

	return pagetemplate, true
}
