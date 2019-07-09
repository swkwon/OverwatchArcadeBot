package owa

import (
	"fmt"
	"strings"
	"testing"
)

func TestTranslateMap(t *testing.T) {
	var prints []string
	for k, v := range translateMap {
		prints = append(prints, fmt.Sprintf("%s@%s", k, v))
	}
	result := strings.Join(prints, "\n")
	Send(result)
}
