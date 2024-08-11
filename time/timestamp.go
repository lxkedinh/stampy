package time

import (
	"strconv"
	"time"
)

func TimeFromSnowflake(snowflake string) time.Time {
	id, _ := strconv.ParseInt(snowflake, 10, 64)
	id = id >> 22
	return time.UnixMilli(id + 1420070400000)
}
