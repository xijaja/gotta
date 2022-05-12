package gottaos

import (
	"runtime"
)

// GottaOS 类型
type GottaOS byte

// WhatIsOS 获取系统信息
func (gos *GottaOS) WhatIsOS() string {
	switch runtime.GOOS {
	case "darwin":
		return "darwin"
	case "linux":
		return "linux"
	case "windows":
		return "windows"
	default:
		return "unknown"
	}
}
