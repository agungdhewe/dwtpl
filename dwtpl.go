package dwtpl

import (
	"io"
	"log"
)

type DeviceType int

const (
	DeviceMobile  DeviceType = 1
	DeviceTablet  DeviceType = 2
	DeviceDesktop DeviceType = 3
)

type TemplateConfig struct {
	Dir    string `yaml:"dir"`
	Cached bool   `yaml:"cached"`
}

var cfg *TemplateConfig
var logger *log.Logger

func New(config *TemplateConfig) {
	cfg = config
	logger = log.New(log.Writer(), "", log.Lmicroseconds|log.Lshortfile)
}

func SetLogOutput(w io.Writer) {
	logger.SetOutput(w)
}

func GetConfig() *TemplateConfig {
	return cfg
}
