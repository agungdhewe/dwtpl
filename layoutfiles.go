package dwtpl

import (
	"fmt"
	"path/filepath"

	"github.com/agungdhewe/dwpath"
)

// GetLayoutFiles retrieves the layout files from a directory based on the configuration file xxx.yml in that directory.
//
// Parameters:
// - dir: the directory path from which to retrieve the layout files.
//
// Returns:
// - files: a map containing the layout files for each device type (DeviceMobile, DeviceTablet, DeviceDesktop).
// - exists: a boolean indicating whether the layout files exist.
// - err: an error if any occurred during the retrieval process.
func (mgr *TemplateManager) GetLayoutFiles(dir string) (files map[DeviceType][]string, exists bool, err error) {

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
			return nil, false, fmt.Errorf("tidak dapat cek file %s", ymllayoutpath)
		} else {
			report_error("file %s tidak ditemukan", ymllayoutpath)
			return nil, false, nil
		}
	}

	// baca konfigurasi
	layoutconfig := &Layout{}
	err = readLayoutConfigYml(ymllayoutpath, layoutconfig)
	if err != nil {
		report_error(err.Error())
		return nil, false, fmt.Errorf("tidak dapat membaca file konfigurasi %s", ymllayoutpath)
	}

	// ambil daftar file sesuai device yang didefinisikan
	files = make(map[DeviceType][]string)
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
