package ip

import (
	"encoding/json"
	"io"
	"net/http"
)

const url = "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip="

// GetIp 获取ip
// 会优先获取 X-Real-IP 和 X-Forwarded-For 头部的值
func GetIp(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	// 如果没有获取到，就使用 RemoteAddr
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

// GetCityByIp 根据 IP 地址获取城市信息
// @param ip string IP 地址
// @return string 城市信息
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}

	// 检查本地 IP 地址
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP" // 本地 IP 地址，返回一个字符串指示它是内网 IP
	}

	// 从 URL 获取响应字节
	resp, err := http.Get(url + ip)
	if err != nil {
		return "" // 如果获取响应时出错，返回空字符串
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 将响应字节转换为 UTF-8 编码的字符串
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "" // 如果读取响应体时出错，返回空字符串
	}
	// 解析 JSON 响应
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "" // 如果解析 JSON 时出错，返回空字符串
	}

	// 检查响应中的代码是否为 0（成功）
	if result["code"].(float64) == 0 {
		city := result["city"].(string) // 获取城市信息
		return city
	} else {
		return ""
	}
}
