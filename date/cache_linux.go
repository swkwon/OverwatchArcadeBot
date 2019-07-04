package date

import (
	"io/ioutil"
	"time"
)

func WriteFile() {
	t := time.Now()
	str := t.Format(time.RFC3339)
	ioutil.WriteFile("/tmp/data.log", []byte(str), 0644)
}

func ReadFile() string {
	if b, e := ioutil.ReadFile("/tmp/data.log"); e == nil {
		return string(b)
	}
	return ""
}
