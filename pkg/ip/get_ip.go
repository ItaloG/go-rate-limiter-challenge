package ip

import (
	"strings"
)

func GetIp(ip string) string {
	if strings.Split(ip, "]")[0] == "[::1" {
		return "127.0.0.1"
	}

	return strings.Split(ip, ":")[0]
}
