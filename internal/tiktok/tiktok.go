package tiktok

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"tiktok-view-bot/internal/signature"
	"time"
)

func NewTikTok(device Device, proxy string) (*TikTok, error) {
	p, err := url.Parse(proxy)
	if err != nil {
		return nil, err
	}

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(p), ForceAttemptHTTP2: true},
		Timeout:   10 * time.Second,
		Jar:       jar,
	}

	return &TikTok{
		client: client,
		device: &device,
	}, nil
}

func (t *TikTok) createParams() string {
	params := map[string]any{
		"iid":                   t.device.InstallId,
		"device_id":             t.device.DeviceId,
		"ac":                    "wifi",
		"channel":               "googleplay",
		"aid":                   1233,
		"app_name":              "musical_ly",
		"version_code":          290304,
		"version_name":          "29.3.4",
		"device_platform":       "android",
		"os":                    "android",
		"ab_version":            "29.3.4",
		"ssmix":                 "a",
		"device_type":           t.device.DeviceInfo.DeviceType,
		"device_brand":          "samsung",
		"language":              "en",
		"os_api":                28,
		"os_version":            9,
		"openudid":              t.device.DeviceInfo.DeviceType,
		"manifest_version_code": 2022903040,
		"resolution":            t.device.DeviceInfo.Resolution,
		"dpi":                   t.device.DeviceInfo.Dpi,
		"update_version_code":   2022903040,
		"app_type":              "normal",
		"sys_region":            t.device.DeviceInfo.SysRegion,
		"mcc_mnc":               50501,
		"timezone_name":         `Asia\Yakutsk`,
		"carrier_region_v2":     505,
		"app_language":          "en",
		"carrier_region":        t.device.DeviceInfo.CarrierRegion,
		"ac2":                   "wifi5g",
		"uoo":                   0,
		"op_region":             t.device.DeviceInfo.OpRegion,
		"timezone_offset":       t.device.DeviceInfo.TimezoneOffset,
		"build_number":          "29.3.4",
		"host_abi":              "arm64-v8a",
		"locale":                t.device.DeviceInfo.Locale,
		"region":                t.device.DeviceInfo.Region,
		"cdid":                  t.device.DeviceInfo.Cdid,
	}

	query := url.Values{}
	for key, value := range params {
		query.Add(key, fmt.Sprintf("%v", value))
	}

	return query.Encode()
}

func (t *TikTok) createHeaders(signature map[string]string, payload string, cookies string) http.Header {
	stub := md5.Sum([]byte(payload))

	return http.Header{
		"content-type":              {"application/x-www-form-urlencoded; charset=UTF-8"},
		"accept-encoding":           {"application/json"},
		"cookie":                    {cookies},
		"passport-sdk-version":      {"19"},
		"sdk-version":               {"2"},
		"user-agent":                {"com.zhiliaoapp.musically/2022903040 (Linux; U; Android 9; en_US; SM-G977N; Build/LMY48Z;tt-ok/3.12.13.1)"},
		"x-argus":                   {"iXezAlBdeoONxmYFYN7vi4wBB39kLVrgUvzvJWzOaRrSr8/Bbj777qHcPUQeKfgfKH7v94B6Cx3b8V6HpCzVKOSobKtOAIYccG+cQ4oVRgcvKo5oLqAD/0P2/VWs2yv5qXgF1sLxEqMg5i2CyDmFEcSUS8D7R08oG0z3RdMv6z5y0UXdYqKMyomJuqCc3xyLQr+RWaaWh7/kN1VfQQd67fJKUDqkfV6lwmssjOhvov1MLJu6q8K/pcfZwwRha9Z/xkSRSD8UYn9uRlQIy+fFUtxu"},
		"x-gorgon":                  {signature["x-gorgon"]},
		"x-khronos":                 {signature["x-khronos"]},
		"x-ladon":                   {"G2GCCV7qqfoOGszXlN4/yP1EViOEmP6wez1kTcy+5ArrTVi1"},
		"x-ss-req-ticket":           {"1682927953076"},
		"x-ss-stub":                 {strings.ToUpper(hex.EncodeToString(stub[:]))},
		"x-tt-dm-status":            {"login=0;ct=1;rt=6"},
		"x-tt-store-region":         {"au"},
		"x-tt-store-region-src":     {"did"},
		"x-vc-bdturing-sdk-version": {"2.3.0.i18n"},
	}
}

func (t *TikTok) View(videoId string) error {
	payload := fmt.Sprintf(`{"item_id": "%v", "play_delta": 1}`, videoId)
	params := t.createParams()
	sig := signature.NewSignature(params, "", "").GetValue()
	headers := t.createHeaders(sig, payload, t.device.Cookie)

	req, err := http.NewRequest(http.MethodPost, "https://api19-core-c-alisg.tiktokv.com/aweme/v1/aweme/stats/?"+params, strings.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header = headers

	res, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// todo: fix empty response.
	fmt.Println(res.StatusCode)
	fmt.Println(string(body))
	return nil
}
