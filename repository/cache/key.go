package cache

import (
	"fmt"
)

const (
	RankKey = "rank"
)

// TaskViewKey 点击数的key
func TaskViewKey(id int64) string {
	return fmt.Sprintf("view:task:%d", id)
}
