package owa

import (
	"fmt"
	"testing"
)

func TestTranslateMap(t *testing.T) {
	for k, v := range translateMap {
		fmt.Println(k, v)
	}
}
