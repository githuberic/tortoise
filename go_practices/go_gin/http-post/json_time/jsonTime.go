package json_time

import "time"

const (
	YYYYMMDD          = "2006-01-02"
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

// JSONTime is time
type JSONTime time.Time

// UnmarshalJSON for JSON Time
func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+YYYYMMDD+`"`, string(data), time.Local)
	*t = JSONTime(now)
	return
}

// MarshalJSON for JSON Time
func (t JSONTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(YYYYMMDD)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, YYYYMMDD)
	b = append(b, '"')
	return b, nil
}

// String for JSON Time
func (t JSONTime) String() string {
	return time.Time(t).Format(YYYYMMDD)
}
