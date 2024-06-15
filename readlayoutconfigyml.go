package dwtpl

import (
	"os"

	"gopkg.in/yaml.v3"
)

// readLayoutConfigYml reads the layout configuration from a YAML file and unmarshals it into the provided Layout struct.
//
// Parameters:
// - filepath: the path to the YAML file containing the layout configuration.
// - layout: a pointer to the Layout struct where the unmarshaled configuration will be stored.
//
// Returns:
// - err: an error if any occurred during the reading or unmarshaling process.
func readLayoutConfigYml(filepath string, layout *Layout) (err error) {
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
