package owa

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTranslateMap(t *testing.T) {
	var prints []string
	for k, v := range translateMap {
		prints = append(prints, fmt.Sprintf("%s@%s", k, v))
	}

	b, e := json.Marshal(prints)
	if e != nil {
		t.Error(e)
	} else {
		Send(string(b))
	}
}
