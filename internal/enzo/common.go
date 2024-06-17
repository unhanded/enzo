package enzo

import "time"

func Now() int64 {
	return time.Now().UnixMilli()
}
