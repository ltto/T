package tp

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type UnixTime time.Time

func (u UnixTime) Value() (driver.Value, error) {
	return time.Time(u), nil
}

func (u *UnixTime) Scan(src interface{}) error {
	*u = UnixTime(src.(time.Time))
	return nil
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, u)
}

func (u UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", time.Time(u).Unix())), nil
}

func (u UnixTime) GormDataType(d gorm.Dialect) string {
	return "timestamp NULL"
}