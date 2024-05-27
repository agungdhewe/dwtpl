package dwtpl

type DeviceType string

const (
	DeviceMobile  DeviceType = "mobile"
	DeviceTablet  DeviceType = "tablet"
	DeviceDesktop DeviceType = "desktop"
)

var mgr *TemplateManager
