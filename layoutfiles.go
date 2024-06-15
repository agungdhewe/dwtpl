package dwtpl

import (
	"fmt"
	"path/filepath"

	"github.com/agungdhewe/dwpath"
)

// GetLayoutFiles retrieves a list of files in a directory to be used for laying out the view
//
// Parameters:
//
//	dir string - the directory path to read layout data from
//
// Returns:
//
//	map[DeviceType][]string - a mapping of DeviceType to a list of file paths
//	bool - indicating if the operation was successful
//	error - an error, if any, that occurred during the operation
func (mgr *TemplateManager) GetLayoutFiles(dir string) (map[DeviceType][]string, bool, error) {
	var err error
	var exists bool

	// siapkan untuk membaca data layout
	basename := filepath.Base(dir)
	ymllayoutfile := fmt.Sprintf("%s.yml", basename)
	ymllayoutpath := filepath.Join(dir, ymllayoutfile)

	// cek apakah direktori ada

	// cek file configurasi yml
	exists, _, err = dwpath.IsFileExists(ymllayoutpath)
	if !exists {
		if err != nil {
			report_error(err.Error())
			return nil, false, fmt.Errorf("file %s tidak ditemukan", ymllayoutpath)
		} else {
			report_log("file %s tidak ditemukan", ymllayoutpath)
			return nil, false, nil
		}
	}

	// baca konfigurasi
	layoutconfig := &Layout{}
	err = readLayoutConfigYml(ymllayoutpath, layoutconfig)
	if err != nil {
		report_error(err.Error())
		return nil, false, fmt.Errorf("tidak dapat baca file konfigurasi %s", ymllayoutpath)
	}

	// ambil daftar file sesuai device yang didefinisikan
	var files = make(map[DeviceType][]string)
	files[DeviceMobile] = layoutconfig.Device.Mobile
	files[DeviceTablet] = layoutconfig.Device.Tablet
	files[DeviceDesktop] = layoutconfig.Device.Desktop

	// tambahkan path pada daftar file
	for _, device := range []DeviceType{DeviceMobile, DeviceTablet, DeviceDesktop} {
		for i, filename := range files[device] {
			files[device][i] = filepath.Join(dir, filename)
		}
	}

	return files, true, nil
}
