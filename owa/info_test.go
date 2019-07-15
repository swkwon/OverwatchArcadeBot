package owa

import (
	"fmt"
	"testing"
)

func TestEmoji(t *testing.T) {
	for k, v := range translateMap{
		fmt.Println(k, v)
	}
}