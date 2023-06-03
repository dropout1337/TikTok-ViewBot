package tiktok

import "net/http"

type TikTok struct {
	client *http.Client
	device *Device
}

type Device struct {
	IsActivated interface{} `json:"is_activated"`
	InstallId   string      `json:"install_id"`
	DeviceId    string      `json:"device_id"`
	NewUser     interface{} `json:"new_user"`
	Cookie      string      `json:"cookie"`
	DeviceInfo  struct {
		Iid                 string `json:"iid"`
		DeviceId            string `json:"device_id"`
		VersionCode         string `json:"version_code"`
		OsVersion           string `json:"os_version"`
		AppName             string `json:"app_name"`
		Channel             string `json:"channel"`
		DevicePlatform      string `json:"device_platform"`
		Aid                 string `json:"aid"`
		Cdid                string `json:"cdid"`
		Openudid            string `json:"openudid"`
		DeviceType          string `json:"device_type"`
		DeviceBrand         string `json:"device_brand"`
		OsApi               int    `json:"os_api"`
		Dpi                 int    `json:"dpi"`
		Ssmix               string `json:"ssmix"`
		Region              string `json:"region"`
		CarrierRegion       string `json:"carrier_region"`
		OpRegion            string `json:"op_region"`
		SysRegion           string `json:"sys_region"`
		AppLanguage         string `json:"app_language"`
		Locale              string `json:"locale"`
		Language            string `json:"language"`
		TimezoneOffset      int    `json:"timezone_offset"`
		ManifestVersionCode string `json:"manifest_version_code"`
		UpdateVersionCode   string `json:"update_version_code"`
		VersionName         string `json:"version_name"`
		AbVersion           string `json:"ab_version"`
		BuildNumber         string `json:"build_number"`
		Ac2                 string `json:"ac2"`
		Ac                  string `json:"ac"`
		Resolution          string `json:"resolution"`
	} `json:"device_info"`
}
