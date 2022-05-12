package times

import (
	"fmt"
	"time"
)

type GottaTimes byte

// HourAgo 计算时间差，支持负数、1.5、0.5 这样的表述，单位为小时
func (gt *GottaTimes) HourAgo(dsc string) string {
	now := time.Now()
	h, _ := time.ParseDuration(dsc)
	h1 := now.Add(1 * h) // 几小时前的时间
	tm := fmt.Sprintf(h1.Format("2006-01-02 15:04:05"))
	return tm
}
