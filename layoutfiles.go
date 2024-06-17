package dwtpl

import (
	"fmt"
	"path/filepath"

	"github.com/agungdhewe/dwpath"
)

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

func (mgr *TemplateManager) GetPageConfig(dir string) (*PageConfig, error) {
	var err error
	var exists bool

	// siapkan untuk membaca data layout
	basename := filepath.Base(dir)
	ymllayoutfile := fmt.Sprintf("%s.yml", basename)
	ymllayoutpath := filepath.Join(dir, ymllayoutfile)

	// cek file configurasi yml
	exists, _, err = dwpath.IsFileExists(ymllayoutpath)
	if !exists {
		if err != nil {
			report_error(err.Error())
		}
		return nil, fmt.Errorf("file %s tidak ditemukan", ymllayoutpath)
	}

	// baca konfigurasi
	pageconfig := &PageConfig{}
	err = readPageConfigYml(ymllayoutpath, pageconfig)
	if err != nil {
		report_error(err.Error())
		return nil, fmt.Errorf("tidak dapat baca file konfigurasi %s", ymllayoutpath)
	}

	var files = make(map[DeviceType][]string)
	files[DeviceMobile] = pageconfig.Device.Mobile
	files[DeviceTablet] = pageconfig.Device.Tablet
	files[DeviceDesktop] = pageconfig.Device.Desktop

	// tambahkan path pada daftar file
	for _, device := range []DeviceType{DeviceMobile, DeviceTablet, DeviceDesktop} {
		for i, filename := range files[device] {
			files[device][i] = filepath.Join(dir, filename)
		}
	}

	pageconfig.layoutfiles = files

	return pageconfig, nil
}
