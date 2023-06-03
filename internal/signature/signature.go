package signature

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NewSignature(params, data, cookies string) *Signature {
	return &Signature{
		Params:  params,
		Data:    data,
		Cookies: cookies,
	}
}

func (s *Signature) hash(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func (s *Signature) getBaseString() string {
	baseStr := s.hash(s.Params)
	if s.Data != "" {
		baseStr = baseStr + s.hash(s.Data)
	} else {
		baseStr = baseStr + strings.Repeat("0", 32)
	}
	if s.Cookies != "" {
		baseStr = baseStr + s.hash(s.Cookies)
	} else {
		baseStr = baseStr + strings.Repeat("0", 32)
	}
	return baseStr
}

func (s *Signature) GetValue() map[string]string {
	return s.encrypt(s.getBaseString())
}

func (s *Signature) encrypt(data string) map[string]string {
	unix := time.Now().Unix()
	length := 0x14
	key := []int{0xDF, 0x77, 0xB9, 0x40, 0xB9, 0x9B, 0x84, 0x83, 0xD1, 0xB9, 0xCB, 0xD1, 0xF7, 0xC2, 0xB9, 0x85, 0xC3, 0xD0, 0xFB, 0xC3}
	paramList := make([]int64, 0)
	for i := 0; i < 12; i += 4 {
		temp := data[8*i : 8*(i+1)]
		for j := 0; j < 4; j++ {
			H, _ := strconv.ParseInt(temp[j*2:(j+1)*2], 16, 64)
			paramList = append(paramList, H)
		}
	}
	paramList = append(paramList, 0x0, 0x6, 0xB, 0x1C)

	paramList = append(paramList, (int64(unix)&0xFF000000)>>24, (int64(unix)&0x00FF0000)>>16, (int64(unix)&0x0000FF00)>>8, (int64(unix)&0x000000FF)>>0)
	eorResultList := make([]int, len(key))
	for i, val := range paramList {
		eorResultList[i] = int(val ^ int64(key[i%len(key)]))
	}
	for i := 0; i < length; i++ {
		C := s.reverse(eorResultList[i])
		D := eorResultList[(i+1)%length]
		E := C ^ D
		F := s.rbitAlgorithm(E)
		H := ((int64(F) ^ 0xFFFFFFFF) ^ int64(length)) & 0xFF
		eorResultList[i] = int(H)
	}
	result := ""
	for _, param := range eorResultList {
		result += s.hexString(param)
	}
	return map[string]string{
		"x-ss-req-ticket": strconv.FormatInt(unix*1000, 10),
		"x-khronos":       strconv.FormatInt(unix, 10),
		"x-gorgon":        "840480e90000" + result,
	}
}

func (s *Signature) rbitAlgorithm(num int) int {
	tmpString := fmt.Sprintf("%08b", num)
	result := ""
	for i := 0; i < 8; i++ {
		result = result + string(tmpString[7-i])
	}
	val, _ := strconv.ParseInt(result, 2, 64)
	return int(val)
}

func (s *Signature) hexString(num int) string {
	tmpString := strconv.FormatInt(int64(num), 16)
	if len(tmpString) < 2 {
		tmpString = "0" + tmpString
	}
	return tmpString
}

func (s *Signature) reverse(num int) int {
	tmpString := s.hexString(num)
	val, _ := strconv.ParseInt(tmpString[1:]+tmpString[:1], 16, 64)
	return int(val)
}
