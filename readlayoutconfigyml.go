package dwtpl

import (
	"os"

	"gopkg.in/yaml.v3"
)

func readLayoutConfigYml(filepath string, layout *Layout) error {
	var err error
	var filedata []byte

	// baca file konfigurasi
	filedata, err = os.ReadFile(filepath)
	if err != nil {
		return err
	}

	// decode file yaml
	err = yaml.Unmarshal(filedata, &layout)
	if err != nil {
		return err
	}

	return nil
}
