package timestamp

import (
	"strconv"
	"time"
)

type Timestamp struct {
	DateString string
	Unix       int64
}

func FromSnowflake(snowflake string) Timestamp {
	id, _ := strconv.ParseInt(snowflake, 10, 64)
	id = id >> 22
	t := time.UnixMilli(id + 1420070400000)
	return Timestamp{
		t.Format(time.RFC822),
		t.Unix(),
	}
}
