package dwtpl

import (
	"fmt"
	"path/filepath"
)

func (mgr *TemplateManager) CachePages(dir string) error {
	mgr.pagesDirLocation = dir

	// baca direktori pagedir
	pattern := filepath.Join(mgr.pagesDirLocation, "*")
	pages, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	for _, pagedir := range pages {
		pagename := filepath.Base(pagedir)
		tpldata, ispage, err := mgr.ParsePageTemplate(pagename, pagedir)
		if err != nil {
			return err
		}

		if ispage {
			mgr.cachedata[pagename] = tpldata
		}
	}

	fmt.Println(mgr.cachedata)

	return nil
}
