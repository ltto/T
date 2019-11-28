package obj

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type UnixTime time.Time

func (t UnixTime) Unix() int64 {
	return time.Time(t).Unix()
}

// MarshalJSON implements json.Marshaler.
func (t UnixTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Time(t).Unix())
	return []byte(stamp), nil
}

func (t UnixTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(t).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(t), nil
}
func (t *UnixTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = UnixTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
