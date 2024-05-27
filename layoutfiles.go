package dwtpl

import (
	"fmt"
	"path/filepath"

	"github.com/agungdhewe/dwpath"
)

// mengambil daftar file-file di suatu direktori yang akan digunakan untuk melayout tampilan
// berdasar file konfigurasi xxx.yml pada direktori tersebut
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
			report_error("ada kesalahan saat cek file %s", ymllayoutpath)
			return nil, false, err
		} else {
			report_error("file %s tidak ditemukan", ymllayoutpath)
			return nil, false, nil
		}
	}

	// baca konfigurasi
	layoutconfig := &Layout{}
	err = readLayoutConfigYml(ymllayoutpath, layoutconfig)
	if err != nil {
		report_error("error saat baca file konfigurasi %s", ymllayoutpath)
		return nil, false, err
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
