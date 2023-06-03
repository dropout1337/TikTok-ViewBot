package main

import (
	"encoding/json"
	"tiktok-view-bot/internal/tiktok"
)

func main() {
	deviceStr := `{"is_activated": "True", "install_id": "7229653131117774598", "device_id": "7229652410409518597", "new_user": "True", "cookie": "store-idc=maliva;store-country-code=pl;store-country-code-src=did;install_id=7229653131117774598;ttreq=1$07560ee124c99b64e8c9e49f65801ca4e3f0740f;", "device_info": {"iid": "7229653131117774598", "device_id": "7229652410409518597", "version_code": "290304", "os_version": "9", "app_name": "musical_ly", "channel": "googleplay", "device_platform": "android", "aid": "1233", "cdid": "8f0faecf-70a0-aa11-b91a-7e3089d770a6", "openudid": "217674f5d1c6087f", "device_type": "rmx2170", "device_brand": "realme", "os_api": 27, "dpi": 480, "ssmix": "a", "region": "PL", "carrier_region": "PL", "op_region": "PL", "sys_region": "PL", "app_language": "en", "locale": "en", "language": "en", "timezone_offset": 32400, "manifest_version_code": "290304", "update_version_code": "290304", "version_name": "29.3.4", "ab_version": "29.3.4", "build_number": "29.3.4", "ac2": "wifi", "ac": "wifi", "resolution": "320*480"}}`
	var device tiktok.Device
	json.Unmarshal([]byte(deviceStr), &device)

	client, _ := tiktok.NewTikTok(device, "http://username:password@domain.com:port")
	client.View("0000000000000000000")
}
