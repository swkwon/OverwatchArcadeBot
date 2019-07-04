package date

import (
	"io/ioutil"
	"time"
)

func WriteFile() {
	t := time.Now()
	str := t.Format(time.RFC3339)
	ioutil.WriteFile("data.log", []byte(str), 0644)
}

func ReadFile() string {
	if b, e := ioutil.ReadFile("data.log"); e == nil {
		return string(b)
	}
	return ""
}
