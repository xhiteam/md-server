package global

import (
	"database/sql/driver"
	"time"

	"fmt"
	"strconv"
)
type LocalTime struct {
	time.Time
}
func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	seconds := t.Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func NormalFormat(datetime string) string {
	t1,_ :=time.Parse(time.RFC3339,datetime)
	return t1.Format("2006-01-02 15:04:05")
}