package dwtpl

import (
	"os"

	"gopkg.in/yaml.v3"
)

// readLayoutConfigYml reads and decodes a YAML layout configuration file.
//
// Parameters:
//
//	filepath string - the path to the YAML file
//	layout *Layout - a pointer to the Layout struct to store the parsed configuration
//
// Return:
//
//	error - an error, if any, that occurred during the operation
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

func readPageConfigYml(filepath string, pageconfig *PageConfig) error {
	var err error
	var filedata []byte

	// baca file konfigurasi
	filedata, err = os.ReadFile(filepath)
	if err != nil {
		return err
	}

	// decode file yaml
	err = yaml.Unmarshal(filedata, &pageconfig)
	if err != nil {
		return err
	}

	return nil
}
