package duration

import (
	"encoding/json"
	str2duration "github.com/xhit/go-str2duration/v2"
	"time"
)

// Duration is a serializable extension of time.Duration
type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' {
		sd := string(b[1 : len(b)-1])
		// use str2duration to support days and weeks
		d.Duration, err = str2duration.ParseDuration(sd)
		return
	}

	var id int64
	id, err = json.Number(b).Int64()
	d.Duration = time.Duration(id)

	return
}

func (d Duration) MarshalJSON() (b []byte, err error) {
	return json.Marshal(d.String())
}
